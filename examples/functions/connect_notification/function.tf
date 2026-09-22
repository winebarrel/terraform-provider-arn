# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/notification/notification-id
output "connect_notification" {
  value = provider::arn::connect_notification("instance-id", "notification-id")
}
