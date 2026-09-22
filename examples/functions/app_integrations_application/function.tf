# arn:aws:app-integrations:ap-northeast-1:111111111111:application/application-id
output "app_integrations_application" {
  value = provider::arn::app_integrations_application("application-id")
}
