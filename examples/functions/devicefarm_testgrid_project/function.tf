# arn:aws:devicefarm:ap-northeast-1:111111111111:testgrid-project:resource-id
output "devicefarm_testgrid_project" {
  value = provider::arn::devicefarm_testgrid_project("resource-id")
}
