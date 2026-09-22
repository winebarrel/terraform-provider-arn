# arn:aws:devicefarm:ap-northeast-1:111111111111:run:resource-id
output "devicefarm_run" {
  value = provider::arn::devicefarm_run("resource-id")
}
