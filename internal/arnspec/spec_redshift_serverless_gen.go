// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: redshift-serverless
// Source: https://servicereference.us-east-1.amazonaws.com/v1/redshift-serverless/redshift-serverless.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "redshift_serverless_endpoint_access", Service: "redshift-serverless", Resource: "endpointAccess", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:managedvpcendpoint/${EndpointAccessId}"},
		{Name: "redshift_serverless_managed_workgroup", Service: "redshift-serverless", Resource: "managed-workgroup", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:managed-workgroup/${ManagedWorkgroupName}"},
		{Name: "redshift_serverless_namespace", Service: "redshift-serverless", Resource: "namespace", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:namespace/${NamespaceId}"},
		{Name: "redshift_serverless_recovery_point", Service: "redshift-serverless", Resource: "recoveryPoint", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:recoverypoint/${RecoveryPointId}"},
		{Name: "redshift_serverless_snapshot", Service: "redshift-serverless", Resource: "snapshot", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:snapshot/${SnapshotId}"},
		{Name: "redshift_serverless_workgroup", Service: "redshift-serverless", Resource: "workgroup", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:workgroup/${WorkgroupId}"},
	})
}
