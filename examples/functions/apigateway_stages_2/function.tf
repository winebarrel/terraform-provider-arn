# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/stages
output "apigateway_stages_2" {
  value = provider::arn::apigateway_stages_2("rest-api-id")
}
