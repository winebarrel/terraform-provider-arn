# arn:aws:events:ap-northeast-1:111111111111:rule/rule-name
output "events_rule_on_default_event_bus" {
  value = provider::arn::events_rule_on_default_event_bus("rule-name")
}
