# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/event-streams/event-stream-name
output "profile_event_streams" {
  value = provider::arn::profile_event_streams("domain-name", "event-stream-name")
}
