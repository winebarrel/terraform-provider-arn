# arn:aws:iot:ap-northeast-1:111111111111:package/package-name/version/version-name
output "iot_packageversion" {
  value = provider::arn::iot_packageversion("package-name", "version-name")
}
