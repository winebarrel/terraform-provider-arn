# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/requestvalidators
output "apigateway_request_validators" {
  value = provider::arn::apigateway_request_validators("rest-api-id")
}
