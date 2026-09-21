package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/winebarrel/terraform-provider-arn/internal/arnconf"
)

// Exported for external (_test package) tests. Only compiled at test time, so
// the package's public API stays unchanged.

// ParseOpts exposes the options-map decoder, whose failure modes (unknown
// key, mutually exclusive keys) are tedious to reach through a full Terraform
// run but are the main thing standing between a typo and a wrong ARN.
func ParseOpts(ctx context.Context, m types.Map) (arnconf.Opts, error) {
	return parseOpts(ctx, m)
}

var DiagSummary = diagSummary
