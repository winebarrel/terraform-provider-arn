// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: pricingplanmanager
// Source: https://servicereference.us-east-1.amazonaws.com/v1/pricingplanmanager/pricingplanmanager.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "pricingplanmanager_subscription", Service: "pricingplanmanager", Resource: "subscription", Template: "arn:${Partition}:pricingplanmanager::${Account}:subscription/${SubscriptionId}"},
	})
}
