# arn:aws:devicefarm:ap-northeast-1:111111111111:devicepool:resource-id
output "devicefarm_devicepool" {
  value = provider::arn::devicefarm_devicepool("resource-id")
}
