# arn:aws:securityhub:ap-northeast-1:111111111111:automation-rule/automation-rule-id
output "securityhub_automation_rule" {
  value = provider::arn::securityhub_automation_rule("automation-rule-id")
}
