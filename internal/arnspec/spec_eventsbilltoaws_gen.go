// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: eventsbilltoaws
// Source: https://servicereference.us-east-1.amazonaws.com/v1/eventsbilltoaws/eventsbilltoaws.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "eventsbilltoaws_approve", Service: "eventsbilltoaws", Resource: "approve", Template: "arn:${Partition}:eventsbilltoaws:${Region}:${Account}:${RelativeId}"},
		{Name: "eventsbilltoaws_info", Service: "eventsbilltoaws", Resource: "info", Template: "arn:${Partition}:eventsbilltoaws:${Region}:${Account}:${RelativeId}"},
	})
}
