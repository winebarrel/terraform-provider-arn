# arn:aws:greengrass:ap-northeast-1:111111111111:deployments:deployment-id
output "greengrass_deployment_2" {
  value = provider::arn::greengrass_deployment_2("deployment-id")
}
