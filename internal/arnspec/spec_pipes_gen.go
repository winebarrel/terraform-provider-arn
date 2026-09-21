// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: pipes
// Source: https://servicereference.us-east-1.amazonaws.com/v1/pipes/pipes.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "pipes_pipe", Service: "pipes", Resource: "pipe", Template: "arn:${Partition}:pipes:${Region}:${Account}:pipe/${Name}"},
	})
}
