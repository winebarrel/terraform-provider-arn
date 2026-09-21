// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: networkflowmonitor
// Source: https://servicereference.us-east-1.amazonaws.com/v1/networkflowmonitor/networkflowmonitor.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "networkflowmonitor_monitor", Service: "networkflowmonitor", Resource: "monitor", Template: "arn:${Partition}:networkflowmonitor:${Region}:${Account}:monitor/${MonitorName}"},
		{Name: "networkflowmonitor_scope", Service: "networkflowmonitor", Resource: "scope", Template: "arn:${Partition}:networkflowmonitor:${Region}:${Account}:scope/${ScopeId}"},
	})
}
