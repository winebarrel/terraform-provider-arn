# arn:aws:networkmonitor:ap-northeast-1:111111111111:monitor/monitor-name
output "networkmonitor_monitor" {
  value = provider::arn::networkmonitor_monitor("monitor-name")
}
