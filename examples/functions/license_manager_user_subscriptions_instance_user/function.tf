# arn:aws:license-manager-user-subscriptions:ap-northeast-1:111111111111:instance-user/instance-user-id
output "license_manager_user_subscriptions_instance_user" {
  value = provider::arn::license_manager_user_subscriptions_instance_user("instance-user-id")
}
