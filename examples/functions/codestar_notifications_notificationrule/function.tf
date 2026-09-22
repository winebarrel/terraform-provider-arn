# arn:aws:codestar-notifications:ap-northeast-1:111111111111:notificationrule/notification-rule-id
output "codestar_notifications_notificationrule" {
  value = provider::arn::codestar_notifications_notificationrule("notification-rule-id")
}
