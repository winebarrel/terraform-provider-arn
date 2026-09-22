# arn:aws:dynamodb:ap-northeast-1:111111111111:table/table-name/import/import-name
output "dynamodb_import" {
  value = provider::arn::dynamodb_import("table-name", "import-name")
}
