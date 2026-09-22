# arn:aws:license-manager-linux-subscriptions:ap-northeast-1:111111111111:subscription-provider/subscription-provider-id
output "license_manager_linux_subscriptions_subscription_provider" {
  value = provider::arn::license_manager_linux_subscriptions_subscription_provider("subscription-provider-id")
}
