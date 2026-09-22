# arn:aws:compute-optimizer::111111111111:automation-rule/rule-id
output "aco_automation_automation_rule" {
  value = provider::arn::aco_automation_automation_rule("rule-id")
}
