// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: application-signals-mcp
// Source: https://servicereference.us-east-1.amazonaws.com/v1/application-signals-mcp/application-signals-mcp.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "application_signals_mcp_mcp_server", Service: "application-signals-mcp", Resource: "mcp-server", Template: "arn:${Partition}:application-signals-mcp:${Region}:${Account}:mcp-server/*"},
	})
}
