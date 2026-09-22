# arn:aws:events:ap-northeast-1:111111111111:event-bus/event-bus-name
output "events_event_bus" {
  value = provider::arn::events_event_bus("event-bus-name")
}
