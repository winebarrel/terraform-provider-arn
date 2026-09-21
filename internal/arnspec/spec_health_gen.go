// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: health
// Source: https://servicereference.us-east-1.amazonaws.com/v1/health/health.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "health_event", Service: "health", Resource: "event", Template: "arn:${Partition}:health:*::event/${Service}/${EventTypeCode}/*"},
	})
}
