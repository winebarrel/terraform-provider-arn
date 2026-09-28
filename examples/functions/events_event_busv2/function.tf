# arn:aws:events:ap-northeast-1:111111111111:event-busv2/event-bus-name/opaque-id
output "events_event_busv2" {
  value = provider::arn::events_event_busv2("event-bus-name", "opaque-id")
}
