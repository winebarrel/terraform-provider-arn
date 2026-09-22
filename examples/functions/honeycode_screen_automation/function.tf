# arn:aws:honeycode:ap-northeast-1:111111111111:screen-automation:workbook/workbook-id/app/app-id/screen/screen-id/automation/automation-id
output "honeycode_screen_automation" {
  value = provider::arn::honeycode_screen_automation("workbook-id", "app-id", "screen-id", "automation-id")
}
