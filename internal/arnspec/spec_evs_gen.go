// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: evs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/evs/evs.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "evs_environment", Service: "evs", Resource: "environment", Template: "arn:${Partition}:evs:${Region}:${Account}:environment/${EnvironmentIdentifier}"},
	})
}
