// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotfleethub
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotfleethub/iotfleethub.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotfleethub_application", Service: "iotfleethub", Resource: "application", Template: "arn:${Partition}:iotfleethub:${Region}:${Account}:application/${ApplicationId}"},
	})
}
