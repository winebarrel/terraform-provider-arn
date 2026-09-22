# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id/environment/environment-id/configuration/configuration-profile-id
output "appconfig_configuration" {
  value = provider::arn::appconfig_configuration("application-id", "environment-id", "configuration-profile-id")
}
