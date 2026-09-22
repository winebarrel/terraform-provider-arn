# arn:aws:codedeploy:ap-northeast-1:111111111111:deploymentconfig:deployment-configuration-name
output "codedeploy_deploymentconfig" {
  value = provider::arn::codedeploy_deploymentconfig("deployment-configuration-name")
}
