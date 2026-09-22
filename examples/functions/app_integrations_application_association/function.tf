# arn:aws:app-integrations:ap-northeast-1:111111111111:application-association/application-id/application-association-id
output "app_integrations_application_association" {
  value = provider::arn::app_integrations_application_association("application-id", "application-association-id")
}
