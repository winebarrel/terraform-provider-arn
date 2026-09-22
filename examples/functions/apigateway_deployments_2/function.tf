# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/deployments
output "apigateway_deployments_2" {
  value = provider::arn::apigateway_deployments_2("rest-api-id")
}
