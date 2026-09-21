// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appstudio
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appstudio/appstudio.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appstudio_application", Service: "appstudio", Resource: "application", Template: "arn:${Partition}:appstudio:${Region}:${Account}:application/${ApplicationId}"},
		{Name: "appstudio_connector", Service: "appstudio", Resource: "connector", Template: "arn:${Partition}:appstudio:${Region}:${Account}:connector/${ConnectionId}"},
		{Name: "appstudio_instance", Service: "appstudio", Resource: "instance", Template: "arn:${Partition}:appstudio:${Region}:${Account}:instance/${InstanceId}"},
	})
}
