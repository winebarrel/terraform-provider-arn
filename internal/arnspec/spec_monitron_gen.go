// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: monitron
// Source: https://servicereference.us-east-1.amazonaws.com/v1/monitron/monitron.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "monitron_project", Service: "monitron", Resource: "project", Template: "arn:${Partition}:monitron:${Region}:${Account}:project/${ResourceId}"},
	})
}
