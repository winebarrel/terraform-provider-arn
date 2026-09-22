# arn:aws:securityhub:ap-northeast-1:111111111111:automation-rulev2/automation-rule-v2-id
output "securityhub_automation_rulev2" {
  value = provider::arn::securityhub_automation_rulev2("automation-rule-v2-id")
}
