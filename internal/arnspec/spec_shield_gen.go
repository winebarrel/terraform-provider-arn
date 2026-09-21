// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: shield
// Source: https://servicereference.us-east-1.amazonaws.com/v1/shield/shield.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "shield_attack", Service: "shield", Resource: "attack", Template: "arn:${Partition}:shield::${Account}:attack/${Id}"},
		{Name: "shield_protection", Service: "shield", Resource: "protection", Template: "arn:${Partition}:shield::${Account}:protection/${Id}"},
		{Name: "shield_protection_group", Service: "shield", Resource: "protection-group", Template: "arn:${Partition}:shield::${Account}:protection-group/${Id}"},
	})
}
