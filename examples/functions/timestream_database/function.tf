# arn:aws:timestream:ap-northeast-1:111111111111:database/database-name
output "timestream_database" {
  value = provider::arn::timestream_database("database-name")
}
