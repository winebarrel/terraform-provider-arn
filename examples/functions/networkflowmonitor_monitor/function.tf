# arn:aws:networkflowmonitor:ap-northeast-1:111111111111:monitor/monitor-name
output "networkflowmonitor_monitor" {
  value = provider::arn::networkflowmonitor_monitor("monitor-name")
}
