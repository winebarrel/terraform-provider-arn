# arn:aws:notifications::111111111111:managed-notification-configuration/category/category/sub-category/subcategory/event/notification-event-id
output "notifications_managed_notification_event" {
  value = provider::arn::notifications_managed_notification_event("category", "subcategory", "notification-event-id")
}
