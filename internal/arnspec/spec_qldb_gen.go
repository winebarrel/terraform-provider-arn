// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: qldb
// Source: https://servicereference.us-east-1.amazonaws.com/v1/qldb/qldb.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "qldb_catalog", Service: "qldb", Resource: "catalog", Template: "arn:${Partition}:qldb:${Region}:${Account}:ledger/${LedgerName}/information_schema/user_tables"},
		{Name: "qldb_ledger", Service: "qldb", Resource: "ledger", Template: "arn:${Partition}:qldb:${Region}:${Account}:ledger/${LedgerName}"},
		{Name: "qldb_stream", Service: "qldb", Resource: "stream", Template: "arn:${Partition}:qldb:${Region}:${Account}:stream/${LedgerName}/${StreamId}"},
		{Name: "qldb_table", Service: "qldb", Resource: "table", Template: "arn:${Partition}:qldb:${Region}:${Account}:ledger/${LedgerName}/table/${TableId}"},
	})
}
