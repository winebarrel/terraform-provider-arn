# arn:aws:events:ap-northeast-1::event-source/event-source-name
output "events_event_source" {
  value = provider::arn::events_event_source("event-source-name")
}
