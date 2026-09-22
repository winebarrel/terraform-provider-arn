# arn:aws:license-manager-user-subscriptions:ap-northeast-1:111111111111:identity-provider/identity-provider-id
output "license_manager_user_subscriptions_identity_provider" {
  value = provider::arn::license_manager_user_subscriptions_identity_provider("identity-provider-id")
}
