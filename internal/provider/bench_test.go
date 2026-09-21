package provider_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/winebarrel/terraform-provider-arn/internal/provider"
)

// BenchmarkProviderSchema models what Terraform asks for when it starts the
// plugin: the full set of functions, each with its name and definition. With
// one function per AWS resource type this is the provider's largest fixed
// cost, so it is the number worth watching.
func BenchmarkProviderSchema(b *testing.B) {
	ctx := context.Background()

	b.ReportAllocs()
	for b.Loop() {
		p := provider.New("bench")()
		pf, ok := p.(interface {
			Functions(context.Context) []func() function.Function
		})
		if !ok {
			b.Fatal("provider does not implement Functions")
		}
		var defs int
		for _, mk := range pf.Functions(ctx) {
			f := mk()
			var meta function.MetadataResponse
			f.Metadata(ctx, function.MetadataRequest{}, &meta)
			var def function.DefinitionResponse
			f.Definition(ctx, function.DefinitionRequest{}, &def)
			defs++
		}
		if defs == 0 {
			b.Fatal("no functions")
		}
	}
}

// BenchmarkFunctionsOnly separates building the 2321 closures from filling in
// their definitions, since only the latter formats strings.
func BenchmarkFunctionsOnly(b *testing.B) {
	ctx := context.Background()
	p := provider.New("bench")().(interface {
		Functions(context.Context) []func() function.Function
	})

	b.ReportAllocs()
	for b.Loop() {
		_ = p.Functions(ctx)
	}
}
