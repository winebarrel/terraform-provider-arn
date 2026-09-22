# arn:aws:notifications:ap-northeast-1:111111111111:configuration/notification-configuration-id/event/notification-event-id
output "notifications_notification_event" {
  value = provider::arn::notifications_notification_event("notification-configuration-id", "notification-event-id")
}
