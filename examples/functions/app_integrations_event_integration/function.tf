# arn:aws:app-integrations:ap-northeast-1:111111111111:event-integration/event-integration-name
output "app_integrations_event_integration" {
  value = provider::arn::app_integrations_event_integration("event-integration-name")
}
