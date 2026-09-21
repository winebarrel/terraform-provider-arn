// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bugbust
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bugbust/bugbust.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bugbust_event", Service: "bugbust", Resource: "Event", Template: "arn:${Partition}:bugbust:${Region}:${Account}:events/${EventId}"},
	})
}
