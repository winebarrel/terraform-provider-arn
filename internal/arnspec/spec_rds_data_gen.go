// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rds-data
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rds-data/rds-data.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rds_data_cluster", Service: "rds-data", Resource: "cluster", Template: "arn:${Partition}:rds:${Region}:${Account}:cluster:${DbClusterInstanceName}"},
	})
}
