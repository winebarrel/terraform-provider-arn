# arn:aws:timestream:ap-northeast-1:111111111111:database/database-name/table/table-name
output "timestream_table" {
  value = provider::arn::timestream_table("database-name", "table-name")
}
