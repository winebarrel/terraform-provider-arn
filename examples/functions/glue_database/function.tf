# arn:aws:glue:ap-northeast-1:111111111111:database/database-name
output "glue_database" {
  value = provider::arn::glue_database("database-name")
}
