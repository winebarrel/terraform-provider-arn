// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sqlworkbench
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sqlworkbench/sqlworkbench.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sqlworkbench_chart", Service: "sqlworkbench", Resource: "chart", Template: "arn:${Partition}:sqlworkbench:${Region}:${Account}:chart/${ResourceId}"},
		{Name: "sqlworkbench_connection", Service: "sqlworkbench", Resource: "connection", Template: "arn:${Partition}:sqlworkbench:${Region}:${Account}:connection/${ResourceId}"},
		{Name: "sqlworkbench_notebook", Service: "sqlworkbench", Resource: "notebook", Template: "arn:${Partition}:sqlworkbench:${Region}:${Account}:notebook/${ResourceId}"},
		{Name: "sqlworkbench_query", Service: "sqlworkbench", Resource: "query", Template: "arn:${Partition}:sqlworkbench:${Region}:${Account}:query/${ResourceId}"},
	})
}
