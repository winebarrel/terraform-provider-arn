# arn:aws:devicefarm:ap-northeast-1:111111111111:sample:resource-id
output "devicefarm_sample" {
  value = provider::arn::devicefarm_sample("resource-id")
}
