# arn:aws:end-user-messaging:ap-northeast-1:111111111111:notify-code-configuration/resource-id
output "end_user_messaging_notify_code_configuration" {
  value = provider::arn::end_user_messaging_notify_code_configuration("resource-id")
}
