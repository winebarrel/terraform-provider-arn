// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: sqs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/sqs/sqs.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "sqs_queue", Service: "sqs", Resource: "queue", Template: "arn:${Partition}:sqs:${Region}:${Account}:${QueueName}"},
	})
}
