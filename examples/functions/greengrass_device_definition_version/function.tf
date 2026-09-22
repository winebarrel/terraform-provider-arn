# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/devices/device-definition-id/versions/version-id
output "greengrass_device_definition_version" {
  value = provider::arn::greengrass_device_definition_version("device-definition-id", "version-id")
}
