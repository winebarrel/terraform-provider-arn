# arn:aws:profile:ap-northeast-1:111111111111:domains/domain-name/event-triggers/event-trigger-name
output "profile_event_triggers" {
  value = provider::arn::profile_event_triggers("domain-name", "event-trigger-name")
}
