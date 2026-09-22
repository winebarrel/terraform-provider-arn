# arn:aws:app-integrations:ap-northeast-1:111111111111:event-integration-association/event-integration-name/resource-id
output "app_integrations_event_integration_association" {
  value = provider::arn::app_integrations_event_integration_association("event-integration-name", "resource-id")
}
