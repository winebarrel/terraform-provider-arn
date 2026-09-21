// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: wickr
// Source: https://servicereference.us-east-1.amazonaws.com/v1/wickr/wickr.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "wickr_network", Service: "wickr", Resource: "network", Template: "arn:${Partition}:wickr:${Region}:${Account}:network/${NetworkId}"},
	})
}
