# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/authorizers
output "apigateway_authorizers_2" {
  value = provider::arn::apigateway_authorizers_2("rest-api-id")
}
