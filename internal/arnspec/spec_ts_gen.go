// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ts
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ts/ts.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ts_execution", Service: "ts", Resource: "execution", Template: "arn:${Partition}:ts::${Account}:execution/${UserId}/${ToolId}/${ExecutionId}"},
		{Name: "ts_tool", Service: "ts", Resource: "tool", Template: "arn:${Partition}:ts::aws:tool/${ToolId}"},
	})
}
