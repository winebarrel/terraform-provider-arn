// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cassandra
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cassandra/cassandra.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cassandra_keyspace", Service: "cassandra", Resource: "keyspace", Template: "arn:${Partition}:cassandra:${Region}:${Account}:/keyspace/${KeyspaceName}/"},
		{Name: "cassandra_stream", Service: "cassandra", Resource: "stream", Template: "arn:${Partition}:cassandra:${Region}:${Account}:/keyspace/${KeyspaceName}/table/${TableName}/stream/${StreamLabel}"},
		{Name: "cassandra_table", Service: "cassandra", Resource: "table", Template: "arn:${Partition}:cassandra:${Region}:${Account}:/keyspace/${KeyspaceName}/table/${TableName}"},
	})
}
