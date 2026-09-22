# arn:aws:medialive:ap-northeast-1:111111111111:eventbridge-rule-template-group:event-bridge-rule-template-group-id
output "medialive_eventbridge_rule_template_group" {
  value = provider::arn::medialive_eventbridge_rule_template_group("event-bridge-rule-template-group-id")
}
