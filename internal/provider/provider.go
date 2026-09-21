// Package provider implements the arn provider: a set of functions that build
// AWS ARNs from a small configuration file, so a configuration does not have
// to interpolate data.aws_caller_identity into every ARN string.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/winebarrel/terraform-provider-arn/internal/arnconf"
	"github.com/winebarrel/terraform-provider-arn/internal/arnspec"
)

var _ provider.ProviderWithFunctions = &ARNProvider{}

type ARNProvider struct {
	version string

	// cache holds the parsed configuration file for the life of the provider
	// process, which Terraform keeps alive for a whole plan or apply. Sharing
	// it across every function means the file is read once, and a malformed
	// file reports the same error everywhere instead of once per call site.
	cache *arnconf.Cache
}

func (p *ARNProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "arn"
	resp.Version = p.version
}

// Schema is empty on purpose. Provider-defined functions cannot read provider
// configuration, so putting account_id or region in a provider block would
// declare settings the functions could never see. The configuration file is
// the only input.
func (p *ARNProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{Attributes: map[string]schema.Attribute{}}
}

func (p *ARNProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

func (p *ARNProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *ARNProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Functions returns one function per generated spec. Terraform asks for this
// list during provider startup, before any configuration is evaluated, which
// is why the specs are generated into the binary rather than fetched from the
// AWS service reference feed at run time.
func (p *ARNProvider) Functions(_ context.Context) []func() function.Function {
	all := arnspec.All()
	out := make([]func() function.Function, 0, len(all))
	for _, s := range all {
		out = append(out, func() function.Function {
			return ARNFunction{spec: s, cache: p.cache}
		})
	}
	return out
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ARNProvider{version: version, cache: &arnconf.Cache{}}
	}
}
