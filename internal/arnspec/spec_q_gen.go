// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: q
// Source: https://servicereference.us-east-1.amazonaws.com/v1/q/q.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "q_plugin", Service: "q", Resource: "plugin", Template: "arn:${Partition}:qdeveloper:${Region}:${Account}:plugin/${Identifier}"},
		{Name: "q_profile", Service: "q", Resource: "profile", Template: "arn:${Partition}:codewhisperer:${Region}:${Account}:profile/${Identifier}"},
	})
}
