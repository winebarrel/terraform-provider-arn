# arn:aws:devicefarm:ap-northeast-1:111111111111:project:resource-id
output "devicefarm_project" {
  value = provider::arn::devicefarm_project("resource-id")
}
