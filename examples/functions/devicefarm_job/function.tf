# arn:aws:devicefarm:ap-northeast-1:111111111111:job:resource-id
output "devicefarm_job" {
  value = provider::arn::devicefarm_job("resource-id")
}
