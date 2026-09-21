// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: payment-cryptography
// Source: https://servicereference.us-east-1.amazonaws.com/v1/payment-cryptography/payment-cryptography.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "payment_cryptography_alias", Service: "payment-cryptography", Resource: "alias", Template: "arn:${Partition}:payment-cryptography:${Region}:${Account}:alias/${Alias}"},
		{Name: "payment_cryptography_approval_team", Service: "payment-cryptography", Resource: "approval-team", Template: "arn:${Partition}:mpa:${Region}:${Account}:approval-team/${ApprovalTeamId}"},
		{Name: "payment_cryptography_key", Service: "payment-cryptography", Resource: "key", Template: "arn:${Partition}:payment-cryptography:${Region}:${Account}:key/${KeyId}"},
	})
}
