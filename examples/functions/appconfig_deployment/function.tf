# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id/environment/environment-id/deployment/deployment-number
output "appconfig_deployment" {
  value = provider::arn::appconfig_deployment("application-id", "environment-id", "deployment-number")
}
