# arn:aws:apigateway:ap-northeast-1::/apis/api-id/deployments/deployment-id
output "apigateway_deployment" {
  value = provider::arn::apigateway_deployment("api-id", "deployment-id")
}
