# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/devices/device-definition-id
output "greengrass_device_definition" {
  value = provider::arn::greengrass_device_definition("device-definition-id")
}
