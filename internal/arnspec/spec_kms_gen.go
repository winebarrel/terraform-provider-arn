// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kms
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kms/kms.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kms_alias", Service: "kms", Resource: "alias", Template: "arn:${Partition}:kms:${Region}:${Account}:alias/${Alias}"},
		{Name: "kms_key", Service: "kms", Resource: "key", Template: "arn:${Partition}:kms:${Region}:${Account}:key/${KeyId}"},
	})
}
