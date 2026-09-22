# arn:aws:greengrass:ap-northeast-1:111111111111:coreDevices:core-device-thing-name
output "greengrass_core_device" {
  value = provider::arn::greengrass_core_device("core-device-thing-name")
}
