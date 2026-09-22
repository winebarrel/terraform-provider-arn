# arn:aws:dynamodb:ap-northeast-1:111111111111:table/table-name/export/export-name
output "dynamodb_export" {
  value = provider::arn::dynamodb_export("table-name", "export-name")
}
