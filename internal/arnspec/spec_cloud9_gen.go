// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloud9
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloud9/cloud9.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloud9_environment", Service: "cloud9", Resource: "environment", Template: "arn:${Partition}:cloud9:${Region}:${Account}:environment:${ResourceId}"},
	})
}
