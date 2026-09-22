# arn:aws:notifications::111111111111:managed-notification-configuration/category/category/sub-category/subcategory
output "notifications_managed_notification_configuration" {
  value = provider::arn::notifications_managed_notification_configuration("category", "subcategory")
}
