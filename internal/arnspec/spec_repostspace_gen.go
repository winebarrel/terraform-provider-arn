// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: repostspace
// Source: https://servicereference.us-east-1.amazonaws.com/v1/repostspace/repostspace.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "repostspace_space", Service: "repostspace", Resource: "space", Template: "arn:${Partition}:repostspace:${Region}:${Account}:space/${ResourceId}"},
	})
}
