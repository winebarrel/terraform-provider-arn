// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: b2bi
// Source: https://servicereference.us-east-1.amazonaws.com/v1/b2bi/b2bi.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "b2bi_capability", Service: "b2bi", Resource: "capability", Template: "arn:${Partition}:b2bi:${Region}:${Account}:capability/${ResourceId}"},
		{Name: "b2bi_partnership", Service: "b2bi", Resource: "partnership", Template: "arn:${Partition}:b2bi:${Region}:${Account}:partnership/${ResourceId}"},
		{Name: "b2bi_profile", Service: "b2bi", Resource: "profile", Template: "arn:${Partition}:b2bi:${Region}:${Account}:profile/${ResourceId}"},
		{Name: "b2bi_transformer", Service: "b2bi", Resource: "transformer", Template: "arn:${Partition}:b2bi:${Region}:${Account}:transformer/${ResourceId}"},
	})
}
