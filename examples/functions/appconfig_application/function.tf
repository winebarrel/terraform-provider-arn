# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id
output "appconfig_application" {
  value = provider::arn::appconfig_application("application-id")
}
