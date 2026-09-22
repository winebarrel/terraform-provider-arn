# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/resources/resource-id/methods/http-method-type/integration
output "apigateway_integration_2" {
  value = provider::arn::apigateway_integration_2("rest-api-id", "resource-id", "http-method-type")
}
