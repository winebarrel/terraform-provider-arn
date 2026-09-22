# arn:aws:glue:ap-northeast-1:111111111111:table/database-name/table-name
output "glue_table" {
  value = provider::arn::glue_table("database-name", "table-name")
}
