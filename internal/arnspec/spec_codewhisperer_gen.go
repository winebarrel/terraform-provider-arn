// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codewhisperer
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codewhisperer/codewhisperer.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codewhisperer_customization", Service: "codewhisperer", Resource: "customization", Template: "arn:${Partition}:codewhisperer:${Region}:${Account}:customization/${Identifier}"},
		{Name: "codewhisperer_profile", Service: "codewhisperer", Resource: "profile", Template: "arn:${Partition}:codewhisperer:${Region}:${Account}:profile/${Identifier}"},
	})
}
