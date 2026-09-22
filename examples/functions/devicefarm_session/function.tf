# arn:aws:devicefarm:ap-northeast-1:111111111111:session:resource-id
output "devicefarm_session" {
  value = provider::arn::devicefarm_session("resource-id")
}
