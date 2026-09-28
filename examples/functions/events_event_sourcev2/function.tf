# arn:aws:events:ap-northeast-1:111111111111:event-sourcev2/source-type/event-source-name/opaque-id
output "events_event_sourcev2" {
  value = provider::arn::events_event_sourcev2("source-type", "event-source-name", "opaque-id")
}
