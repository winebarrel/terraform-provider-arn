# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id
output "apigateway_rest_api" {
  value = provider::arn::apigateway_rest_api("rest-api-id")
}
