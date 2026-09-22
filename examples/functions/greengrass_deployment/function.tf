# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/groups/group-id/deployments/deployment-id
output "greengrass_deployment" {
  value = provider::arn::greengrass_deployment("group-id", "deployment-id")
}
