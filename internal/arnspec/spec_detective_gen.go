// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: detective
// Source: https://servicereference.us-east-1.amazonaws.com/v1/detective/detective.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "detective_graph", Service: "detective", Resource: "Graph", Template: "arn:${Partition}:detective:${Region}:${Account}:graph:${ResourceId}"},
	})
}
