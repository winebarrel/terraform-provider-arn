// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sdb
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sdb/sdb.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sdb_domain", Service: "sdb", Resource: "domain", Template: "arn:${Partition}:sdb:${Region}:${Account}:domain/${DomainName}"},
		{Name: "sdb_export", Service: "sdb", Resource: "export", Template: "arn:${Partition}:sdb:${Region}:${Account}:domain/${DomainName}/export/${ExportUUID}"},
	})
}
