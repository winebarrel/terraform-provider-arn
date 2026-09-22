# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id/configurationprofile/configuration-profile-id
output "appconfig_configurationprofile" {
  value = provider::arn::appconfig_configurationprofile("application-id", "configuration-profile-id")
}
