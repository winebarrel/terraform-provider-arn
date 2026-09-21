// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: dsql
// Source: https://servicereference.us-east-1.amazonaws.com/v1/dsql/dsql.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "dsql_cluster", Service: "dsql", Resource: "Cluster", Template: "arn:${Partition}:dsql:${Region}:${Account}:cluster/${Identifier}"},
		{Name: "dsql_stream", Service: "dsql", Resource: "Stream", Template: "arn:${Partition}:dsql:${Region}:${Account}:cluster/${ClusterId}/stream/${StreamId}"},
	})
}
