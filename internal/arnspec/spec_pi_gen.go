// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: pi
// Source: https://servicereference.us-east-1.amazonaws.com/v1/pi/pi.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "pi_metric_resource", Service: "pi", Resource: "metric-resource", Template: "arn:${Partition}:pi:${Region}:${Account}:metrics/${ServiceType}/${Identifier}"},
		{Name: "pi_perf_reports_resource", Service: "pi", Resource: "perf-reports-resource", Template: "arn:${Partition}:pi:${Region}:${Account}:perf-reports/${ServiceType}/${Identifier}/${ReportId}"},
	})
}
