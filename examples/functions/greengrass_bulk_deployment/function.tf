# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/bulk/deployments/bulk-deployment-id
output "greengrass_bulk_deployment" {
  value = provider::arn::greengrass_bulk_deployment("bulk-deployment-id")
}
