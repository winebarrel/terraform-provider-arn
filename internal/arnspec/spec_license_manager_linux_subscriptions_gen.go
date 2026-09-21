// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: license-manager-linux-subscriptions
// Source: https://servicereference.us-east-1.amazonaws.com/v1/license-manager-linux-subscriptions/license-manager-linux-subscriptions.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "license_manager_linux_subscriptions_subscription_provider", Service: "license-manager-linux-subscriptions", Resource: "subscription-provider", Template: "arn:${Partition}:license-manager-linux-subscriptions:${Region}:${Account}:subscription-provider/${SubscriptionProviderId}"},
	})
}
