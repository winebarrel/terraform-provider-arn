# arn:aws:sagemaker:ap-northeast-1:111111111111:edge-deployment/edge-deployment-plan-name
output "sagemaker_edge_deployment_plan" {
  value = provider::arn::sagemaker_edge_deployment_plan("edge-deployment-plan-name")
}
