// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: internetmonitor
// Source: https://servicereference.us-east-1.amazonaws.com/v1/internetmonitor/internetmonitor.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "internetmonitor_health_event", Service: "internetmonitor", Resource: "HealthEvent", Template: "arn:${Partition}:internetmonitor:${Region}:${Account}:monitor/${MonitorName}/health-event/${EventId}"},
		{Name: "internetmonitor_internet_event", Service: "internetmonitor", Resource: "InternetEvent", Template: "arn:${Partition}:internetmonitor::${Account}:internet-event/${InternetEventId}"},
		{Name: "internetmonitor_monitor", Service: "internetmonitor", Resource: "Monitor", Template: "arn:${Partition}:internetmonitor:${Region}:${Account}:monitor/${MonitorName}"},
	})
}
