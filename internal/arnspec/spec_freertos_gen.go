// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: freertos
// Source: https://servicereference.us-east-1.amazonaws.com/v1/freertos/freertos.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "freertos_configuration", Service: "freertos", Resource: "configuration", Template: "arn:${Partition}:freertos:${Region}:${Account}:configuration/${ConfigurationName}"},
		{Name: "freertos_subscription", Service: "freertos", Resource: "subscription", Template: "arn:${Partition}:freertos:${Region}:${Account}:subscription/${SubscriptionID}"},
	})
}
