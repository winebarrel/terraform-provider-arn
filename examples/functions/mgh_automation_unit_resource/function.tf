# arn:aws:mgh:ap-northeast-1:111111111111:automation-unit/automation-unit-id
output "mgh_automation_unit_resource" {
  value = provider::arn::mgh_automation_unit_resource("automation-unit-id")
}
