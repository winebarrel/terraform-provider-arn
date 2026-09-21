// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: s3tables
// Source: https://servicereference.us-east-1.amazonaws.com/v1/s3tables/s3tables.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "s3tables_table", Service: "s3tables", Resource: "Table", Template: "arn:${Partition}:s3tables:${Region}:${Account}:bucket/${TableBucketName}/table/${TableID}"},
		{Name: "s3tables_table_bucket", Service: "s3tables", Resource: "TableBucket", Template: "arn:${Partition}:s3tables:${Region}:${Account}:bucket/${TableBucketName}"},
	})
}
