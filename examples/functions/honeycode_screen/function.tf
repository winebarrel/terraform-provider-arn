# arn:aws:honeycode:ap-northeast-1:111111111111:screen:workbook/workbook-id/app/app-id/screen/screen-id
output "honeycode_screen" {
  value = provider::arn::honeycode_screen("workbook-id", "app-id", "screen-id")
}
