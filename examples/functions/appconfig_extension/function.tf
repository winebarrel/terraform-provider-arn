# arn:aws:appconfig:ap-northeast-1:111111111111:extension/extension-id/extension-version-number
output "appconfig_extension" {
  value = provider::arn::appconfig_extension("extension-id", "extension-version-number")
}
