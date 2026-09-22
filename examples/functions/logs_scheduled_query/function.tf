# arn:aws:logs:ap-northeast-1:111111111111:scheduled-query:scheduled-query-id
output "logs_scheduled_query" {
  value = provider::arn::logs_scheduled_query("scheduled-query-id")
}
