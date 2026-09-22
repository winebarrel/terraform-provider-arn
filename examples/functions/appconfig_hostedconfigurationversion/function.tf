# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id/configurationprofile/configuration-profile-id
output "appconfig_hostedconfigurationversion" {
  value = provider::arn::appconfig_hostedconfigurationversion("application-id", "configuration-profile-id")
}
