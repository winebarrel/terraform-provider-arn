package arnspec

import (
	"fmt"
	"sort"
	"sync"
)

// specs accumulates every generated spec. Each generated file contributes its
// own service through an init, so the package needs no index listing them and
// a service that leaves the AWS feed leaves by having its file deleted.
var specs []Spec

// register is called from the generated files. Package initialization is
// single-threaded, so the append needs no lock.
func register(s []Spec) {
	specs = append(specs, s...)
}

var (
	loadOnce sync.Once
	all      []*Spec
	byName   map[string]*Spec
)

// load derives the parsed form of every spec on first use. It runs after all
// the generated inits, which is why it is not an init of its own: init order
// within a package is by filename, and depending on it would make the
// registry sensitive to what the generator happens to name a file.
//
// A parse failure here is a bug in cmd/gen, which validates every template
// before writing a file, so there is no runtime path that can recover from it
// and panicking is the honest response.
func load() {
	loadOnce.Do(func() {
		all = make([]*Spec, 0, len(specs))
		byName = make(map[string]*Spec, len(specs))

		for i := range specs {
			s := &specs[i]
			if err := Parse(s); err != nil {
				panic(fmt.Sprintf("arnspec: %v", err))
			}
			all = append(all, s)
			byName[s.Name] = s
		}
		// The generated files arrive one service at a time, so order by
		// function name here rather than relying on the order they were read.
		sort.Slice(all, func(i, j int) bool { return all[i].Name < all[j].Name })
	})
}

// All returns every generated spec, ordered by function name.
func All() []*Spec {
	load()
	return all
}

// Lookup returns the spec for a function name.
func Lookup(name string) (*Spec, bool) {
	load()
	s, ok := byName[name]
	return s, ok
}
