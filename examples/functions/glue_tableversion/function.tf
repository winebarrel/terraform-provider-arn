# arn:aws:glue:ap-northeast-1:111111111111:tableVersion/database-name/table-name/table-version-name
output "glue_tableversion" {
  value = provider::arn::glue_tableversion("database-name", "table-name", "table-version-name")
}
