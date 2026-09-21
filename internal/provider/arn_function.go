package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/winebarrel/terraform-provider-arn/internal/arnconf"
	"github.com/winebarrel/terraform-provider-arn/internal/arnspec"
)

var _ function.Function = ARNFunction{}

// Option keys accepted by the trailing options argument.
const (
	optAccount   = "account"
	optAccountID = "account_id"
	optRegion    = "region"
	optPartition = "partition"
)

// optionsParamName is the name of the trailing variadic parameter. It is not
// a Terraform keyword argument, only a label that shows up in documentation
// and in argument-position errors.
const optionsParamName = "options"

// ARNFunction is one generated function. Every function in this provider is
// this same type with a different spec, so there are 2000-odd instances and
// exactly one implementation.
type ARNFunction struct {
	spec  *arnspec.Spec
	cache *arnconf.Cache
}

func (f ARNFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = f.spec.Name
}

func (f ARNFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	params := make([]function.Parameter, 0, len(f.spec.Args))
	for i, a := range f.spec.Args {
		params = append(params, function.StringParameter{
			Name:                a,
			MarkdownDescription: fmt.Sprintf("Value for `${%s}` in the ARN template.", f.spec.ArgRaw(i)),
		})
	}

	resp.Definition = function.Definition{
		Summary:             fmt.Sprintf("Builds an %s %s ARN", f.spec.Service, f.spec.Resource),
		MarkdownDescription: f.markdownDescription(),
		Parameters:          params,
		VariadicParameter: function.MapParameter{
			Name:                optionsParamName,
			ElementType:         types.StringType,
			MarkdownDescription: optionsDoc,
		},
		Return: function.StringReturn{},
	}
}

const optionsDoc = "Optional overrides: `account` (the name of an `account` block in the configuration file), " +
	"`account_id` (a literal account id), `region` and `partition`. " +
	"At most one options map may be given."

func (f ARNFunction) markdownDescription() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Builds the ARN of an `%s` `%s`:\n\n```\n%s\n```\n\n",
		f.spec.Service, f.spec.Resource, f.spec.Template)
	fmt.Fprintf(&b, "`${Partition}`, `${Region}` and `${Account}` come from the configuration file "+
		"(`.arn.hcl`, or the path in `ARN_CONFIG`).")
	if len(f.spec.Args) > 0 {
		b.WriteString(" The remaining placeholders are the arguments, in template order.")
	}
	return b.String()
}

func (f ARNFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	// Arguments.Get insists on a target for every parameter including the
	// variadic one, so the targets are built to match the spec exactly.
	args := make([]string, len(f.spec.Args))
	targets := make([]any, 0, len(args)+1)
	for i := range args {
		targets = append(targets, &args[i])
	}
	var opts []types.Map
	targets = append(targets, &opts)

	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, targets...))
	if resp.Error != nil {
		return
	}

	// The variadic parameter is the only way to express an optional argument,
	// so a caller can pass several maps. Only one is meaningful; silently
	// using the first would hide a mistake.
	if len(opts) > 1 {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(
			int64(len(args)+1),
			fmt.Sprintf("%s takes at most one options map, got %d", f.spec.Name, len(opts)),
		))
		return
	}

	var o arnconf.Opts
	if len(opts) == 1 {
		var err error
		o, err = parseOpts(ctx, opts[0])
		if err != nil {
			resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(int64(len(args)), err.Error()))
			return
		}
	}

	cfg, err := f.cache.Get(arnconf.DefaultPath())
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(
			fmt.Sprintf("read %s: %v", arnconf.DefaultPath(), err)))
		return
	}

	values, err := cfg.Resolve(o)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewArgumentFuncError(int64(len(args)), err.Error()))
		return
	}

	arn, err := f.spec.Build(arnspec.Values{
		Partition: values.Partition,
		Region:    values.Region,
		AccountID: values.AccountID,
	}, args)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(resp.Error, function.NewFuncError(err.Error()))
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, arn))
}

// parseOpts converts the options map into arnconf.Opts. An unrecognized key
// is an error rather than an ignored entry: a typo like "acccount" would
// otherwise fall back to the default account and produce a valid-looking ARN
// pointing at the wrong place.
func parseOpts(ctx context.Context, m types.Map) (arnconf.Opts, error) {
	if m.IsNull() || m.IsUnknown() {
		return arnconf.Opts{}, nil
	}

	raw := map[string]types.String{}
	if diags := m.ElementsAs(ctx, &raw, false); diags.HasError() {
		return arnconf.Opts{}, fmt.Errorf("invalid options map: %s", diagSummary(diags))
	}

	var o arnconf.Opts
	var unknown []string
	for k, v := range raw {
		if v.IsNull() || v.IsUnknown() {
			continue
		}
		switch k {
		case optAccount:
			o.Account = v.ValueString()
		case optAccountID:
			o.AccountID = v.ValueString()
		case optRegion:
			o.Region = v.ValueString()
		case optPartition:
			o.Partition = v.ValueString()
		default:
			unknown = append(unknown, k)
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return arnconf.Opts{}, fmt.Errorf("unknown option(s) %s: valid options are %s, %s, %s, %s",
			strings.Join(unknown, ", "), optAccount, optAccountID, optPartition, optRegion)
	}
	if o.Account != "" && o.AccountID != "" {
		return arnconf.Opts{}, fmt.Errorf("%s and %s are mutually exclusive", optAccount, optAccountID)
	}
	return o, nil
}
