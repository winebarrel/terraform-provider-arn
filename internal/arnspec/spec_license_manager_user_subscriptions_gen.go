// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: license-manager-user-subscriptions
// Source: https://servicereference.us-east-1.amazonaws.com/v1/license-manager-user-subscriptions/license-manager-user-subscriptions.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "license_manager_user_subscriptions_identity_provider", Service: "license-manager-user-subscriptions", Resource: "identity-provider", Template: "arn:${Partition}:license-manager-user-subscriptions:${Region}:${Account}:identity-provider/${IdentityProviderId}"},
		{Name: "license_manager_user_subscriptions_instance_user", Service: "license-manager-user-subscriptions", Resource: "instance-user", Template: "arn:${Partition}:license-manager-user-subscriptions:${Region}:${Account}:instance-user/${InstanceUserId}"},
		{Name: "license_manager_user_subscriptions_license_server_endpoint", Service: "license-manager-user-subscriptions", Resource: "license-server-endpoint", Template: "arn:${Partition}:license-manager-user-subscriptions:${Region}:${Account}:license-server-endpoint/${LicenseServerEndpointId}"},
		{Name: "license_manager_user_subscriptions_product_subscription", Service: "license-manager-user-subscriptions", Resource: "product-subscription", Template: "arn:${Partition}:license-manager-user-subscriptions:${Region}:${Account}:product-subscription/${ProductSubscriptionId}"},
	})
}
