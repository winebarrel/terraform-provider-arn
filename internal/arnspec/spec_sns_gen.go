// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sns
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sns/sns.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sns_topic", Service: "sns", Resource: "topic", Template: "arn:${Partition}:sns:${Region}:${Account}:${TopicName}"},
	})
}
