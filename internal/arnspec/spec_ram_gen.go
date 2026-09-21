// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ram
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ram/ram.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ram_customer_managed_permission", Service: "ram", Resource: "customer-managed-permission", Template: "arn:${Partition}:ram:${Region}:${Account}:permission/${ResourcePath}"},
		{Name: "ram_permission", Service: "ram", Resource: "permission", Template: "arn:${Partition}:ram::${Account}:permission/${ResourcePath}"},
		{Name: "ram_resource_share", Service: "ram", Resource: "resource-share", Template: "arn:${Partition}:ram:${Region}:${Account}:resource-share/${ResourcePath}"},
		{Name: "ram_resource_share_invitation", Service: "ram", Resource: "resource-share-invitation", Template: "arn:${Partition}:ram:${Region}:${Account}:resource-share-invitation/${ResourcePath}"},
	})
}
