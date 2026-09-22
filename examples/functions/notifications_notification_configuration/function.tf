# arn:aws:notifications::111111111111:configuration/notification-configuration-id
output "notifications_notification_configuration" {
  value = provider::arn::notifications_notification_configuration("notification-configuration-id")
}
