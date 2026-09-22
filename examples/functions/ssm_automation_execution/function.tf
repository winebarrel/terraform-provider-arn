# arn:aws:ssm:ap-northeast-1:111111111111:automation-execution/automation-execution-id
output "ssm_automation_execution" {
  value = provider::arn::ssm_automation_execution("automation-execution-id")
}
