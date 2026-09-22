# arn:aws:notifications::111111111111:configuration/notification-configuration-id/rule/event-rule-id
output "notifications_event_rule" {
  value = provider::arn::notifications_event_rule("notification-configuration-id", "event-rule-id")
}
