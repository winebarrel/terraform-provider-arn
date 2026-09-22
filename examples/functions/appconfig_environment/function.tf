# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id/environment/environment-id
output "appconfig_environment" {
  value = provider::arn::appconfig_environment("application-id", "environment-id")
}
