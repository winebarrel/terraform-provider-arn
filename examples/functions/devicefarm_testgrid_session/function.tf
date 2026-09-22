# arn:aws:devicefarm:ap-northeast-1:111111111111:testgrid-session:resource-id
output "devicefarm_testgrid_session" {
  value = provider::arn::devicefarm_testgrid_session("resource-id")
}
