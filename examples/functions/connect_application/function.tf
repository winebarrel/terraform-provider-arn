# arn:aws:app-integrations:ap-northeast-1:111111111111:application/application-id
output "connect_application" {
  value = provider::arn::connect_application("application-id")
}
