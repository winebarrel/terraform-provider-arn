// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: account
// Source: https://servicereference.us-east-1.amazonaws.com/v1/account/account.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "account_account", Service: "account", Resource: "account", Template: "arn:${Partition}:account::${Account}:account"},
		{Name: "account_account_in_organization", Service: "account", Resource: "accountInOrganization", Template: "arn:${Partition}:account::${ManagementAccountId}:account/o-${OrganizationId}/${MemberAccountId}"},
	})
}
