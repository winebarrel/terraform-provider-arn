// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: timestream
// Source: https://servicereference.us-east-1.amazonaws.com/v1/timestream/timestream.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "timestream_database", Service: "timestream", Resource: "database", Template: "arn:${Partition}:timestream:${Region}:${Account}:database/${DatabaseName}"},
		{Name: "timestream_scheduled_query", Service: "timestream", Resource: "scheduled-query", Template: "arn:${Partition}:timestream:${Region}:${Account}:scheduled-query/${ScheduledQueryName}"},
		{Name: "timestream_table", Service: "timestream", Resource: "table", Template: "arn:${Partition}:timestream:${Region}:${Account}:database/${DatabaseName}/table/${TableName}"},
	})
}
