# arn:aws:events:ap-northeast-1:111111111111:rule/event-bus-name/rule-name
output "events_rule_on_custom_event_bus" {
  value = provider::arn::events_rule_on_custom_event_bus("event-bus-name", "rule-name")
}
