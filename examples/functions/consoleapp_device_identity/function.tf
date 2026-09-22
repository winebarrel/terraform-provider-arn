# arn:aws:consoleapp::111111111111:device/device-id/identity/identity-id
output "consoleapp_device_identity" {
  value = provider::arn::consoleapp_device_identity("device-id", "identity-id")
}
