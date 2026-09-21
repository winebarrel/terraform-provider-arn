// Command gen turns the AWS service reference feed into
// internal/arnspec/spec_gen.go.
//
// It runs from `go generate`, never from the provider itself: the provider
// has to be able to declare its functions at startup with no network, and the
// set of functions has to be stable across runs so a plan does not change
// because AWS published a new resource type this morning.
//
//	go run ./cmd/gen -out internal/arnspec/spec_gen.go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/winebarrel/terraform-provider-arn/internal/arnspec"
)

const (
	defaultEndpoint = "https://servicereference.us-east-1.amazonaws.com"
	fetchWorkers    = 16
	httpTimeout     = 30 * time.Second
)

type indexEntry struct {
	Service string `json:"service"`
	URL     string `json:"url"`
}

type serviceDoc struct {
	Name      string `json:"Name"`
	Resources []struct {
		Name       string   `json:"Name"`
		ARNFormats []string `json:"ARNFormats"`
	} `json:"Resources"`
}

type entry struct {
	Name     string
	Service  string
	Resource string
	Template string
}

func main() {
	endpoint := flag.String("endpoint", defaultEndpoint, "AWS service reference endpoint")
	out := flag.String("out", "internal/arnspec/spec_gen.go", "output file")
	flag.Parse()

	client := &http.Client{Timeout: httpTimeout}

	var index []indexEntry
	if err := getJSON(client, strings.TrimRight(*endpoint, "/")+"/", &index); err != nil {
		log.Fatalf("fetch index: %v", err)
	}
	sort.Slice(index, func(i, j int) bool { return index[i].Service < index[j].Service })
	log.Printf("index: %d services", len(index))

	docs := fetchAll(client, index)

	entries, err := build(docs)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("functions: %d", len(entries))

	if err := os.WriteFile(*out, render(entries, *endpoint), 0o644); err != nil { //nolint:gosec
		log.Fatalf("write %s: %v", *out, err)
	}
	if err := exec.Command("gofmt", "-w", *out).Run(); err != nil {
		log.Fatalf("gofmt %s: %v", *out, err)
	}
}

func fetchAll(client *http.Client, index []indexEntry) []serviceDoc {
	docs := make([]serviceDoc, len(index))
	var wg sync.WaitGroup
	sem := make(chan struct{}, fetchWorkers)
	var failed sync.Map

	for i, e := range index {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			if err := getJSON(client, e.URL, &docs[i]); err != nil {
				failed.Store(e.Service, err)
			}
		}()
	}
	wg.Wait()

	// A partial fetch would silently drop functions from the provider, so
	// stop rather than generate a file that is quietly missing a service.
	var bad []string
	failed.Range(func(k, v any) bool {
		bad = append(bad, fmt.Sprintf("%s: %v", k, v))
		return true
	})
	if len(bad) > 0 {
		sort.Strings(bad)
		log.Fatalf("fetch failed for %d service(s):\n  %s", len(bad), strings.Join(bad, "\n  "))
	}
	return docs
}

func build(docs []serviceDoc) ([]entry, error) {
	var entries []entry
	seen := map[string]entry{}

	for _, d := range docs {
		if d.Name == "" {
			continue
		}
		svc := arnspec.SnakeCase(d.Name)
		for _, r := range d.Resources {
			if len(r.ARNFormats) == 0 {
				continue
			}
			base := svc + "_" + arnspec.SnakeCase(r.Name)
			for i, tmpl := range r.ARNFormats {
				// A handful of resource types carry more than one format,
				// always because AWS kept an older API's ARN shape alongside
				// the current one (apigateway's /restapis next to /apis, lex
				// v1 next to v2). There is no name in the feed to tell them
				// apart, so the extras take a numeric suffix; the generated
				// documentation shows the template each one builds.
				name := base
				if i > 0 {
					name = fmt.Sprintf("%s_%d", base, i+1)
				}
				e := entry{Name: name, Service: d.Name, Resource: r.Name, Template: tmpl}
				if prev, dup := seen[name]; dup {
					return nil, fmt.Errorf("duplicate function name %q: %s/%s and %s/%s",
						name, prev.Service, prev.Resource, e.Service, e.Resource)
				}
				seen[name] = e
				entries = append(entries, e)
			}
		}
	}

	sort.Slice(entries, func(i, j int) bool { return entries[i].Name < entries[j].Name })

	// Every template must parse, and a function whose name or arguments would
	// not be a valid Terraform identifier must not reach the generated file.
	for _, e := range entries {
		s := arnspec.Spec{Name: e.Name, Service: e.Service, Resource: e.Resource, Template: e.Template}
		if err := arnspec.Parse(&s); err != nil {
			return nil, err
		}
		if !validIdent(s.Name) {
			return nil, fmt.Errorf("%s/%s: invalid function name %q", e.Service, e.Resource, s.Name)
		}
		for _, a := range s.Args {
			if !validIdent(a) {
				return nil, fmt.Errorf("%s: invalid argument name %q", s.Name, a)
			}
		}
	}
	return entries, nil
}

// validIdent reports whether s is usable as a Terraform function or parameter
// name: lowercase letters, digits and underscores, not starting with a digit.
func validIdent(s string) bool {
	if s == "" || (s[0] >= '0' && s[0] <= '9') {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < 'a' || c > 'z') && (c < '0' || c > '9') && c != '_' {
			return false
		}
	}
	return true
}

func render(entries []entry, endpoint string) []byte {
	var b strings.Builder
	fmt.Fprintf(&b, `// Code generated by cmd/gen; DO NOT EDIT.
//
// Source: %s
// Functions: %d

package arnspec

var specs = []Spec{
`, endpoint, len(entries))

	for _, e := range entries {
		fmt.Fprintf(&b, "\t{Name: %q, Service: %q, Resource: %q, Template: %q},\n",
			e.Name, e.Service, e.Resource, e.Template)
	}
	b.WriteString("}\n")
	return []byte(b.String())
}

func getJSON(client *http.Client, url string, out any) error {
	resp, err := client.Get(url) //nolint:noctx
	if err != nil {
		return err
	}
	defer resp.Body.Close() //nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		_, _ = io.Copy(io.Discard, resp.Body)
		return fmt.Errorf("http %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
