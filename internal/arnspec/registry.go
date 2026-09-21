package arnspec

import "fmt"

// All is every generated spec, ordered by function name.
var All []*Spec

var byName map[string]*Spec

// init derives the parsed form of each generated spec once, at process start.
//
// A parse failure here is a bug in cmd/gen, which validates every template
// before writing spec_gen.go, so there is no runtime path that can recover
// from it and panicking is the honest response.
func init() {
	All = make([]*Spec, 0, len(specs))
	byName = make(map[string]*Spec, len(specs))

	for i := range specs {
		s := &specs[i]
		if err := Parse(s); err != nil {
			panic(fmt.Sprintf("arnspec: %v", err))
		}
		All = append(All, s)
		byName[s.Name] = s
	}
}

// Lookup returns the spec for a function name.
func Lookup(name string) (*Spec, bool) {
	s, ok := byName[name]
	return s, ok
}
