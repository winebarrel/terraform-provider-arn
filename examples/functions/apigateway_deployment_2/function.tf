# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/deployments/deployment-id
output "apigateway_deployment_2" {
  value = provider::arn::apigateway_deployment_2("rest-api-id", "deployment-id")
}
