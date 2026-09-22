# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/resources/resource-id/methods/http-method-type/responses/status-code
output "apigateway_method_response" {
  value = provider::arn::apigateway_method_response("rest-api-id", "resource-id", "http-method-type", "status-code")
}
