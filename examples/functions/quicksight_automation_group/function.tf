# arn:aws:quicksight:ap-northeast-1:111111111111:automation-group/resource-id
output "quicksight_automation_group" {
  value = provider::arn::quicksight_automation_group("resource-id")
}
