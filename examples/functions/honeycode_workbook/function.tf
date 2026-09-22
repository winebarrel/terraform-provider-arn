# arn:aws:honeycode:ap-northeast-1:111111111111:workbook:workbook/workbook-id
output "honeycode_workbook" {
  value = provider::arn::honeycode_workbook("workbook-id")
}
