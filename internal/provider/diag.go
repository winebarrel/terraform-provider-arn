package provider

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// diagSummary flattens framework diagnostics into a single line, for the
// cases where a conversion failure has to be folded into a function error
// (function.FuncError carries text, not diagnostics).
func diagSummary(diags diag.Diagnostics) string {
	parts := make([]string, 0, len(diags))
	for _, d := range diags {
		if detail := d.Detail(); detail != "" {
			parts = append(parts, d.Summary()+": "+detail)
			continue
		}
		parts = append(parts, d.Summary())
	}
	return strings.Join(parts, "; ")
}
