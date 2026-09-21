// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bedrock-websearch
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bedrock-websearch/bedrock-websearch.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bedrock_websearch_tool", Service: "bedrock-websearch", Resource: "tool", Template: "arn:${Partition}:bedrock-websearch:${Region}:aws:tool/${ToolName}"},
	})
}
