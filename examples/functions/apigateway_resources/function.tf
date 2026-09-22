# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/resources
output "apigateway_resources" {
  value = provider::arn::apigateway_resources("rest-api-id")
}
