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
	"path/filepath"
	"sort"
	"strconv"
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
	dir := flag.String("dir", "internal/arnspec", "output directory for the spec tables")
	examples := flag.String("examples", "examples/functions", "output directory for the documentation examples")
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

	written, err := write(*dir, entries, *endpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("spec files: %d", written)

	written, err = writeExamples(*examples, entries)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("examples: %d", written)
}

// Values the examples are written against. They match the configuration file
// the README opens with, so a reader moving between the two sees the same
// account and region.
const (
	exampleAccountID = "111111111111"
	exampleRegion    = "ap-northeast-1"
	examplePartition = "aws"
)

// writeExamples emits examples/functions/<name>/function.tf for every
// function, which is where tfplugindocs looks for the example it puts on each
// function's page. Writing them by hand would mean 2321 files kept in step
// with the templates they call.
//
// Each example calls the function with its own argument names as values, so
// the ARN in the comment above it shows which part of the result each
// argument became.
func writeExamples(dir string, entries []entry) (int, error) {
	if err := os.MkdirAll(dir, 0o755); err != nil { //nolint:gosec
		return 0, err
	}

	want := make(map[string]struct{}, len(entries))
	for _, e := range entries {
		want[e.Name] = struct{}{}
	}

	// A function that leaves the feed leaves by having its directory removed.
	existing, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}
	for _, d := range existing {
		if !d.IsDir() {
			continue
		}
		if _, keep := want[d.Name()]; !keep {
			if err := os.RemoveAll(filepath.Join(dir, d.Name())); err != nil {
				return 0, fmt.Errorf("remove %s: %w", d.Name(), err)
			}
		}
	}

	for _, e := range entries {
		s := arnspec.Spec{Name: e.Name, Service: e.Service, Resource: e.Resource, Template: e.Template}
		if err := arnspec.Parse(&s); err != nil {
			return 0, err
		}

		// The argument name doubles as its value, so the ARN below shows
		// where each one lands.
		args := make([]string, len(s.Args))
		for i, a := range s.Args {
			args[i] = strings.ReplaceAll(a, "_", "-")
		}
		arn, err := s.Build(arnspec.Values{
			Partition: examplePartition,
			Region:    exampleRegion,
			AccountID: exampleAccountID,
		}, args)
		if err != nil {
			return 0, fmt.Errorf("%s: %w", e.Name, err)
		}

		quoted := make([]string, len(args))
		for i, a := range args {
			quoted[i] = strconv.Quote(a)
		}

		// No generated-file header: tfplugindocs embeds this verbatim, so
		// anything here also lands on the function's published page, where
		// it means nothing to the reader. The page already declares itself
		// generated in its own front matter.
		var b strings.Builder
		fmt.Fprintf(&b, "# %s\n", arn)
		fmt.Fprintf(&b, "output %q {\n  value = provider::arn::%s(%s)\n}\n",
			e.Name, e.Name, strings.Join(quoted, ", "))

		sub := filepath.Join(dir, e.Name)
		if err := os.MkdirAll(sub, 0o755); err != nil { //nolint:gosec
			return 0, err
		}
		if err := os.WriteFile(filepath.Join(sub, "function.tf"), []byte(b.String()), 0o644); err != nil { //nolint:gosec
			return 0, err
		}
	}
	return len(entries), nil
}

// write emits one file per service and removes the files of services that are
// no longer in the feed. Splitting by service keeps each file small enough to
// read and makes a diff after regeneration show which services actually
// changed, instead of one 2000-line file moving around.
func write(dir string, entries []entry, endpoint string) (int, error) {
	byService := map[string][]entry{}
	for _, e := range entries {
		byService[e.Service] = append(byService[e.Service], e)
	}

	// Two services must not map to the same file. They cannot today, since
	// the feed's service names are unique and already lowercase, but the
	// filename goes through SnakeCase and a future name with a different
	// separator could collide.
	paths := map[string]string{}
	for svc := range byService {
		p := filepath.Join(dir, "spec_"+arnspec.SnakeCase(svc)+"_gen.go")
		if prev, dup := paths[p]; dup {
			return 0, fmt.Errorf("services %q and %q both map to %s", prev, svc, p)
		}
		paths[p] = svc
	}

	stale, err := filepath.Glob(filepath.Join(dir, "spec_*_gen.go"))
	if err != nil {
		return 0, err
	}
	for _, p := range stale {
		if _, keep := paths[p]; !keep {
			if err := os.Remove(p); err != nil {
				return 0, fmt.Errorf("remove %s: %w", p, err)
			}
		}
	}

	for p, svc := range paths {
		if err := os.WriteFile(p, render(svc, byService[svc], endpoint), 0o644); err != nil { //nolint:gosec
			return 0, fmt.Errorf("write %s: %w", p, err)
		}
	}
	if err := exec.Command("gofmt", "-w", dir).Run(); err != nil {
		return 0, fmt.Errorf("gofmt %s: %w", dir, err)
	}
	return len(paths), nil
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

func render(service string, entries []entry, endpoint string) []byte {
	var b strings.Builder
	fmt.Fprintf(&b, `// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: %s
// Source: %s/v1/%s/%s.json
// Functions: %d
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
`, service, endpoint, service, service, len(entries))

	for _, e := range entries {
		fmt.Fprintf(&b, "\t\t{Name: %q, Service: %q, Resource: %q, Template: %q},\n",
			e.Name, e.Service, e.Resource, e.Template)
	}
	b.WriteString("\t})\n}\n")
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
