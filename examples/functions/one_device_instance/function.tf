# arn:aws:one:ap-northeast-1:111111111111:device-instance/device-instance-id
output "one_device_instance" {
  value = provider::arn::one_device_instance("device-instance-id")
}
