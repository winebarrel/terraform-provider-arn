# arn:aws:internetmonitor:ap-northeast-1:111111111111:monitor/monitor-name/health-event/event-id
output "internetmonitor_health_event" {
  value = provider::arn::internetmonitor_health_event("monitor-name", "event-id")
}
