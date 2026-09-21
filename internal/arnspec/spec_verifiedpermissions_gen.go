// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: verifiedpermissions
// Source: https://servicereference.us-east-1.amazonaws.com/v1/verifiedpermissions/verifiedpermissions.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "verifiedpermissions_policy_store", Service: "verifiedpermissions", Resource: "policy-store", Template: "arn:${Partition}:verifiedpermissions::${Account}:policy-store/${PolicyStoreId}"},
		{Name: "verifiedpermissions_policy_store_alias", Service: "verifiedpermissions", Resource: "policy-store-alias", Template: "arn:${Partition}:verifiedpermissions:${Region}:${Account}:policy-store-alias/${AliasName}"},
	})
}
