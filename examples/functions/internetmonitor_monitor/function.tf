# arn:aws:internetmonitor:ap-northeast-1:111111111111:monitor/monitor-name
output "internetmonitor_monitor" {
  value = provider::arn::internetmonitor_monitor("monitor-name")
}
