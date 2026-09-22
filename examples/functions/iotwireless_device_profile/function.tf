# arn:aws:iotwireless:ap-northeast-1:111111111111:DeviceProfile/device-profile-id
output "iotwireless_device_profile" {
  value = provider::arn::iotwireless_device_profile("device-profile-id")
}
