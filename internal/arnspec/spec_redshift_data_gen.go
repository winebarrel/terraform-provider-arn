// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: redshift-data
// Source: https://servicereference.us-east-1.amazonaws.com/v1/redshift-data/redshift-data.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "redshift_data_cluster", Service: "redshift-data", Resource: "cluster", Template: "arn:${Partition}:redshift:${Region}:${Account}:cluster:${ClusterName}"},
		{Name: "redshift_data_managed_workgroup", Service: "redshift-data", Resource: "managed-workgroup", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:managed-workgroup/${ManagedWorkgroupId}"},
		{Name: "redshift_data_workgroup", Service: "redshift-data", Resource: "workgroup", Template: "arn:${Partition}:redshift-serverless:${Region}:${Account}:workgroup/${WorkgroupId}"},
	})
}
