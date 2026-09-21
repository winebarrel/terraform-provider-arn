// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: vendor-insights
// Source: https://servicereference.us-east-1.amazonaws.com/v1/vendor-insights/vendor-insights.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "vendor_insights_data_source", Service: "vendor-insights", Resource: "DataSource", Template: "arn:${Partition}:vendor-insights:::data-source:${ResourceId}"},
		{Name: "vendor_insights_security_profile", Service: "vendor-insights", Resource: "SecurityProfile", Template: "arn:${Partition}:vendor-insights:::security-profile:${ResourceId}"},
	})
}
