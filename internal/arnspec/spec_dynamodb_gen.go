// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: dynamodb
// Source: https://servicereference.us-east-1.amazonaws.com/v1/dynamodb/dynamodb.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "dynamodb_backup", Service: "dynamodb", Resource: "backup", Template: "arn:${Partition}:dynamodb:${Region}:${Account}:table/${TableName}/backup/${BackupName}"},
		{Name: "dynamodb_export", Service: "dynamodb", Resource: "export", Template: "arn:${Partition}:dynamodb:${Region}:${Account}:table/${TableName}/export/${ExportName}"},
		{Name: "dynamodb_global_table", Service: "dynamodb", Resource: "global-table", Template: "arn:${Partition}:dynamodb::${Account}:global-table/${GlobalTableName}"},
		{Name: "dynamodb_import", Service: "dynamodb", Resource: "import", Template: "arn:${Partition}:dynamodb:${Region}:${Account}:table/${TableName}/import/${ImportName}"},
		{Name: "dynamodb_index", Service: "dynamodb", Resource: "index", Template: "arn:${Partition}:dynamodb:${Region}:${Account}:table/${TableName}/index/${IndexName}"},
		{Name: "dynamodb_stream", Service: "dynamodb", Resource: "stream", Template: "arn:${Partition}:dynamodb:${Region}:${Account}:table/${TableName}/stream/${StreamLabel}"},
		{Name: "dynamodb_table", Service: "dynamodb", Resource: "table", Template: "arn:${Partition}:dynamodb:${Region}:${Account}:table/${TableName}"},
	})
}
