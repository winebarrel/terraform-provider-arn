// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rum
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rum/rum.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rum_app_monitor_resource", Service: "rum", Resource: "AppMonitorResource", Template: "arn:${Partition}:rum:${Region}:${Account}:appmonitor/${Name}"},
	})
}
