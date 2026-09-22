# arn:aws:app-integrations:ap-northeast-1:111111111111:data-integration/data-integration-id
output "app_integrations_data_integration" {
  value = provider::arn::app_integrations_data_integration("data-integration-id")
}
