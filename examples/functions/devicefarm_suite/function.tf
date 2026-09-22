# arn:aws:devicefarm:ap-northeast-1:111111111111:suite:resource-id
output "devicefarm_suite" {
  value = provider::arn::devicefarm_suite("resource-id")
}
