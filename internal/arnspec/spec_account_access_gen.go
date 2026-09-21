// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: account-access
// Source: https://servicereference.us-east-1.amazonaws.com/v1/account-access/account-access.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "account_access_application", Service: "account-access", Resource: "application", Template: "arn:${Partition}:account-access:${Region}:${Account}:application/${ResourceId}"},
	})
}
