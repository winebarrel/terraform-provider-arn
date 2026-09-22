# arn:aws:quicksight:ap-northeast-1:111111111111:automation-group/automation-group-id/automation/resource-id
output "quicksight_automation" {
  value = provider::arn::quicksight_automation("automation-group-id", "resource-id")
}
