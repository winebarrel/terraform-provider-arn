// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: consoleapp
// Source: https://servicereference.us-east-1.amazonaws.com/v1/consoleapp/consoleapp.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "consoleapp_device_identity", Service: "consoleapp", Resource: "DeviceIdentity", Template: "arn:${Partition}:consoleapp::${Account}:device/${DeviceId}/identity/${IdentityId}"},
	})
}
