# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/authorizers/authorizer-id
output "apigateway_authorizer_2" {
  value = provider::arn::apigateway_authorizer_2("rest-api-id", "authorizer-id")
}
