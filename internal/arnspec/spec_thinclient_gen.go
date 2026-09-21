// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: thinclient
// Source: https://servicereference.us-east-1.amazonaws.com/v1/thinclient/thinclient.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "thinclient_device", Service: "thinclient", Resource: "device", Template: "arn:${Partition}:thinclient:${Region}:${Account}:device/${DeviceId}"},
		{Name: "thinclient_environment", Service: "thinclient", Resource: "environment", Template: "arn:${Partition}:thinclient:${Region}:${Account}:environment/${EnvironmentId}"},
		{Name: "thinclient_softwareset", Service: "thinclient", Resource: "softwareset", Template: "arn:${Partition}:thinclient:${Region}:${Account}:softwareset/${SoftwareSetId}"},
	})
}
