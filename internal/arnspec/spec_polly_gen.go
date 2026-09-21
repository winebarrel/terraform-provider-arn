// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: polly
// Source: https://servicereference.us-east-1.amazonaws.com/v1/polly/polly.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "polly_lexicon", Service: "polly", Resource: "lexicon", Template: "arn:${Partition}:polly:${Region}:${Account}:lexicon/${LexiconName}"},
	})
}
