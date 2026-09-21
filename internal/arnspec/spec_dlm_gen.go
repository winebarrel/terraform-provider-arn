// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: dlm
// Source: https://servicereference.us-east-1.amazonaws.com/v1/dlm/dlm.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "dlm_policy", Service: "dlm", Resource: "policy", Template: "arn:${Partition}:dlm:${Region}:${Account}:policy/${ResourceName}"},
	})
}
