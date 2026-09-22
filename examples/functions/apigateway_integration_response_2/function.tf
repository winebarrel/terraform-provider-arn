# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/resources/resource-id/methods/http-method-type/integration/responses/status-code
output "apigateway_integration_response_2" {
  value = provider::arn::apigateway_integration_response_2("rest-api-id", "resource-id", "http-method-type", "status-code")
}
