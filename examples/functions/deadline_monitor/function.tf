# arn:aws:deadline:ap-northeast-1:111111111111:monitor/monitor-id
output "deadline_monitor" {
  value = provider::arn::deadline_monitor("monitor-id")
}
