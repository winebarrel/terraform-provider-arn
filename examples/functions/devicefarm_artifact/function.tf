# arn:aws:devicefarm:ap-northeast-1:111111111111:artifact:resource-id
output "devicefarm_artifact" {
  value = provider::arn::devicefarm_artifact("resource-id")
}
