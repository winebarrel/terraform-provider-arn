# arn:aws:honeycode:ap-northeast-1:111111111111:table:workbook/workbook-id/table/table-id
output "honeycode_table" {
  value = provider::arn::honeycode_table("workbook-id", "table-id")
}
