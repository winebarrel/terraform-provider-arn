// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotmanagedintegrations
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotmanagedintegrations/iotmanagedintegrations.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotmanagedintegrations_account_association", Service: "iotmanagedintegrations", Resource: "account-association", Template: "arn:${Partition}:iotmanagedintegrations:${Region}:${Account}:account-association/${AccountAssociationId}"},
		{Name: "iotmanagedintegrations_credential_locker", Service: "iotmanagedintegrations", Resource: "credential-locker", Template: "arn:${Partition}:iotmanagedintegrations:${Region}:${Account}:credential-locker/${Identifier}"},
		{Name: "iotmanagedintegrations_managed_thing", Service: "iotmanagedintegrations", Resource: "managed-thing", Template: "arn:${Partition}:iotmanagedintegrations:${Region}:${Account}:managed-thing/${Identifier}"},
		{Name: "iotmanagedintegrations_ota_task", Service: "iotmanagedintegrations", Resource: "ota-task", Template: "arn:${Partition}:iotmanagedintegrations:${Region}:${Account}:ota-task/${Identifier}"},
		{Name: "iotmanagedintegrations_provisioning_profile", Service: "iotmanagedintegrations", Resource: "provisioning-profile", Template: "arn:${Partition}:iotmanagedintegrations:${Region}:${Account}:provisioning-profile/${Identifier}"},
	})
}
