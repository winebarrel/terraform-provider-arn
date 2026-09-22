# arn:aws:medialive:ap-northeast-1:111111111111:eventbridge-rule-template:event-bridge-rule-template-id
output "medialive_eventbridge_rule_template" {
  value = provider::arn::medialive_eventbridge_rule_template("event-bridge-rule-template-id")
}
