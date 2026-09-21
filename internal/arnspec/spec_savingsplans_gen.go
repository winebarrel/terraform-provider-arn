// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: savingsplans
// Source: https://servicereference.us-east-1.amazonaws.com/v1/savingsplans/savingsplans.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "savingsplans_savingsplan", Service: "savingsplans", Resource: "savingsplan", Template: "arn:${Partition}:savingsplans::${Account}:savingsplan/${ResourceId}"},
	})
}
