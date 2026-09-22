# arn:aws:appconfig:ap-northeast-1:111111111111:extensionassociation/extension-association-id
output "appconfig_extensionassociation" {
  value = provider::arn::appconfig_extensionassociation("extension-association-id")
}
