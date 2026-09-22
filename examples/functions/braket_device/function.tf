# arn:aws:braket:*:*:device/device-type/provider/device-id
output "braket_device" {
  value = provider::arn::braket_device("device-type", "provider", "device-id")
}
