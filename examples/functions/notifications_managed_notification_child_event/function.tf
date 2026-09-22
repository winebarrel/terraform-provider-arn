# arn:aws:notifications::111111111111:managed-notification-configuration/category/category/sub-category/subcategory/event/notification-event-id/child-event/notification-child-event-id
output "notifications_managed_notification_child_event" {
  value = provider::arn::notifications_managed_notification_child_event("category", "subcategory", "notification-event-id", "notification-child-event-id")
}
