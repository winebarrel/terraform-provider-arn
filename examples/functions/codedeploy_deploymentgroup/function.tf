# arn:aws:codedeploy:ap-northeast-1:111111111111:deploymentgroup:application-name/deployment-group-name
output "codedeploy_deploymentgroup" {
  value = provider::arn::codedeploy_deploymentgroup("application-name", "deployment-group-name")
}
