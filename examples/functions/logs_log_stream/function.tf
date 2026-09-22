# arn:aws:logs:ap-northeast-1:111111111111:log-group:log-group-name:log-stream:log-stream-name
output "logs_log_stream" {
  value = provider::arn::logs_log_stream("log-group-name", "log-stream-name")
}
