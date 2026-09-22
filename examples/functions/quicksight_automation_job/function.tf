# arn:aws:quicksight:ap-northeast-1:111111111111:automation-group/automation-group-id/automation/automation-id/job/resource-id
output "quicksight_automation_job" {
  value = provider::arn::quicksight_automation_job("automation-group-id", "automation-id", "resource-id")
}
