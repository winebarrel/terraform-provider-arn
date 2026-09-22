# arn:aws:devicefarm:ap-northeast-1:111111111111:upload:resource-id
output "devicefarm_upload" {
  value = provider::arn::devicefarm_upload("resource-id")
}
