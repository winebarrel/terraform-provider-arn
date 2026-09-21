// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: identitystore
// Source: https://servicereference.us-east-1.amazonaws.com/v1/identitystore/identitystore.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "identitystore_all_group_memberships", Service: "identitystore", Resource: "AllGroupMemberships", Template: "arn:${Partition}:identitystore:::membership/*"},
		{Name: "identitystore_all_groups", Service: "identitystore", Resource: "AllGroups", Template: "arn:${Partition}:identitystore:::group/*"},
		{Name: "identitystore_all_users", Service: "identitystore", Resource: "AllUsers", Template: "arn:${Partition}:identitystore:::user/*"},
		{Name: "identitystore_group", Service: "identitystore", Resource: "Group", Template: "arn:${Partition}:identitystore:::group/${GroupId}"},
		{Name: "identitystore_group_membership", Service: "identitystore", Resource: "GroupMembership", Template: "arn:${Partition}:identitystore:::membership/${MembershipId}"},
		{Name: "identitystore_identitystore", Service: "identitystore", Resource: "Identitystore", Template: "arn:${Partition}:identitystore::${Account}:identitystore/${IdentityStoreId}"},
		{Name: "identitystore_user", Service: "identitystore", Resource: "User", Template: "arn:${Partition}:identitystore:::user/${UserId}"},
	})
}
