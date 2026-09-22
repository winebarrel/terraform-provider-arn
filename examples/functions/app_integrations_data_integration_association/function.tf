# arn:aws:app-integrations:ap-northeast-1:111111111111:data-integration-association/data-integration-id/resource-id
output "app_integrations_data_integration_association" {
  value = provider::arn::app_integrations_data_integration_association("data-integration-id", "resource-id")
}
