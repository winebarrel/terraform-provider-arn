# arn:aws:chime:ap-northeast-1:111111111111:app-instance/app-instance-id/bot/app-instance-bot-id
output "chime_app_instance_bot" {
  value = provider::arn::chime_app_instance_bot("app-instance-id", "app-instance-bot-id")
}
