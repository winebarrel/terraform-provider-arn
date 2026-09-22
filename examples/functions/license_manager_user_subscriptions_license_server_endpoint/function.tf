# arn:aws:license-manager-user-subscriptions:ap-northeast-1:111111111111:license-server-endpoint/license-server-endpoint-id
output "license_manager_user_subscriptions_license_server_endpoint" {
  value = provider::arn::license_manager_user_subscriptions_license_server_endpoint("license-server-endpoint-id")
}
