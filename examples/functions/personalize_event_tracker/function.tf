# arn:aws:personalize:ap-northeast-1:111111111111:event-tracker/resource-id
output "personalize_event_tracker" {
  value = provider::arn::personalize_event_tracker("resource-id")
}
