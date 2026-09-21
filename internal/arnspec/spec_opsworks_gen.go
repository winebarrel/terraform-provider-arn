// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: opsworks
// Source: https://servicereference.us-east-1.amazonaws.com/v1/opsworks/opsworks.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "opsworks_stack", Service: "opsworks", Resource: "stack", Template: "arn:${Partition}:opsworks:${Region}:${Account}:stack/${StackId}/"},
	})
}
