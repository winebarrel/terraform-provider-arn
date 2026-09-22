# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/gatewayresponses
output "apigateway_gateway_responses" {
  value = provider::arn::apigateway_gateway_responses("rest-api-id")
}
