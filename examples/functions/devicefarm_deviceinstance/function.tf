# arn:aws:devicefarm:ap-northeast-1::deviceinstance:resource-id
output "devicefarm_deviceinstance" {
  value = provider::arn::devicefarm_deviceinstance("resource-id")
}
