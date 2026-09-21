package provider_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/terraform-provider-arn/internal/arnconf"
	"github.com/winebarrel/terraform-provider-arn/internal/provider"
)

func optsMap(t *testing.T, kv map[string]attr.Value) types.Map {
	t.Helper()
	m, diags := types.MapValue(types.StringType, kv)
	require.False(t, diags.HasError(), "%v", diags)
	return m
}

func TestParseOpts(t *testing.T) {
	ctx := context.Background()

	got, err := provider.ParseOpts(ctx, optsMap(t, map[string]attr.Value{
		"account":   types.StringValue("prod"),
		"region":    types.StringValue("us-east-1"),
		"partition": types.StringValue("aws-cn"),
	}))
	require.NoError(t, err)
	assert.Equal(t, arnconf.Opts{Account: "prod", Region: "us-east-1", Partition: "aws-cn"}, got)

	got, err = provider.ParseOpts(ctx, optsMap(t, map[string]attr.Value{
		"account_id": types.StringValue("999999999999"),
	}))
	require.NoError(t, err)
	assert.Equal(t, arnconf.Opts{AccountID: "999999999999"}, got)
}

func TestParseOptsEmptyAndNull(t *testing.T) {
	ctx := context.Background()

	got, err := provider.ParseOpts(ctx, types.MapNull(types.StringType))
	require.NoError(t, err)
	assert.Equal(t, arnconf.Opts{}, got)

	got, err = provider.ParseOpts(ctx, types.MapUnknown(types.StringType))
	require.NoError(t, err)
	assert.Equal(t, arnconf.Opts{}, got)

	// A null element is treated as absent, not as an empty override. Writing
	// { region = null } is how a configuration says "use the default".
	got, err = provider.ParseOpts(ctx, optsMap(t, map[string]attr.Value{
		"account": types.StringValue("prod"),
		"region":  types.StringNull(),
	}))
	require.NoError(t, err)
	assert.Equal(t, arnconf.Opts{Account: "prod"}, got)
}

func TestParseOptsRejectsUnknownKeys(t *testing.T) {
	_, err := provider.ParseOpts(context.Background(), optsMap(t, map[string]attr.Value{
		"acccount": types.StringValue("prod"),
		"regoin":   types.StringValue("us-east-1"),
	}))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown option(s) acccount, regoin")
	assert.Contains(t, err.Error(), "valid options are account, account_id, partition, region")
}

func TestParseOptsRejectsAccountAndAccountID(t *testing.T) {
	_, err := provider.ParseOpts(context.Background(), optsMap(t, map[string]attr.Value{
		"account":    types.StringValue("prod"),
		"account_id": types.StringValue("9"),
	}))
	require.ErrorContains(t, err, "mutually exclusive")
}

func TestDiagSummary(t *testing.T) {
	assert.Empty(t, provider.DiagSummary(nil))

	diags := diag.Diagnostics{
		diag.NewErrorDiagnostic("summary only", ""),
		diag.NewErrorDiagnostic("with detail", "the detail"),
	}
	assert.Equal(t, "summary only; with detail: the detail", provider.DiagSummary(diags))
}

func TestParseOptsRejectsEmptyValues(t *testing.T) {
	_, err := provider.ParseOpts(context.Background(), optsMap(t, map[string]attr.Value{
		"account": types.StringValue(""),
		"region":  types.StringValue(""),
	}))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "option(s) account, region set to an empty string")
}
