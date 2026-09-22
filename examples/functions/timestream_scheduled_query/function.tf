# arn:aws:timestream:ap-northeast-1:111111111111:scheduled-query/scheduled-query-name
output "timestream_scheduled_query" {
  value = provider::arn::timestream_scheduled_query("scheduled-query-name")
}
