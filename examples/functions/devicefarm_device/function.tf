# arn:aws:devicefarm:ap-northeast-1::device:resource-id
output "devicefarm_device" {
  value = provider::arn::devicefarm_device("resource-id")
}
