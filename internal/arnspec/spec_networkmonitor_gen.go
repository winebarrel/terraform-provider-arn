// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: networkmonitor
// Source: https://servicereference.us-east-1.amazonaws.com/v1/networkmonitor/networkmonitor.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "networkmonitor_monitor", Service: "networkmonitor", Resource: "monitor", Template: "arn:${Partition}:networkmonitor:${Region}:${Account}:monitor/${MonitorName}"},
		{Name: "networkmonitor_probe", Service: "networkmonitor", Resource: "probe", Template: "arn:${Partition}:networkmonitor:${Region}:${Account}:probe/${ProbeId}"},
	})
}
