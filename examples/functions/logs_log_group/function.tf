# arn:aws:logs:ap-northeast-1:111111111111:log-group:log-group-name
output "logs_log_group" {
  value = provider::arn::logs_log_group("log-group-name")
}
