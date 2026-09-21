// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bcm-dashboards
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bcm-dashboards/bcm-dashboards.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bcm_dashboards_dashboard", Service: "bcm-dashboards", Resource: "dashboard", Template: "arn:${Partition}:bcm-dashboards::${Account}:dashboard/${DashboardName}"},
		{Name: "bcm_dashboards_scheduled_report", Service: "bcm-dashboards", Resource: "scheduled-report", Template: "arn:${Partition}:bcm-dashboards::${Account}:scheduled-report/${ScheduledReportName}"},
	})
}
