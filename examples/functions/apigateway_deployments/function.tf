# arn:aws:apigateway:ap-northeast-1::/apis/api-id/deployments
output "apigateway_deployments" {
  value = provider::arn::apigateway_deployments("api-id")
}
