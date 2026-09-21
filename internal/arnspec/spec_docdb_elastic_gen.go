// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: docdb-elastic
// Source: https://servicereference.us-east-1.amazonaws.com/v1/docdb-elastic/docdb-elastic.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "docdb_elastic_cluster", Service: "docdb-elastic", Resource: "cluster", Template: "arn:${Partition}:docdb-elastic:${Region}:${Account}:cluster/${ResourceId}"},
		{Name: "docdb_elastic_cluster_snapshot", Service: "docdb-elastic", Resource: "cluster-snapshot", Template: "arn:${Partition}:docdb-elastic:${Region}:${Account}:cluster-snapshot/${ResourceId}"},
	})
}
