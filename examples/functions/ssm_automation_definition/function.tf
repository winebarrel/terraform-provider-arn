# arn:aws:ssm:ap-northeast-1:111111111111:automation-definition/automation-definition-name:version-id
output "ssm_automation_definition" {
  value = provider::arn::ssm_automation_definition("automation-definition-name", "version-id")
}
