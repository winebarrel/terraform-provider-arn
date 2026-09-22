# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/gatewayresponses/response-type
output "apigateway_gateway_response" {
  value = provider::arn::apigateway_gateway_response("rest-api-id", "response-type")
}
