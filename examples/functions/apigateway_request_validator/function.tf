# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/requestvalidators/request-validator-id
output "apigateway_request_validator" {
  value = provider::arn::apigateway_request_validator("rest-api-id", "request-validator-id")
}
