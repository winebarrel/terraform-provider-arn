// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: glacier
// Source: https://servicereference.us-east-1.amazonaws.com/v1/glacier/glacier.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "glacier_vault", Service: "glacier", Resource: "vault", Template: "arn:${Partition}:glacier:${Region}:${Account}:vaults/${VaultName}"},
	})
}
