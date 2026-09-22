# arn:aws:robomaker:ap-northeast-1:111111111111:deployment-fleet/fleet-name/created-on-epoch
output "robomaker_deployment_fleet" {
  value = provider::arn::robomaker_deployment_fleet("fleet-name", "created-on-epoch")
}
