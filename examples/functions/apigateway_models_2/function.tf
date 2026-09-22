# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/models
output "apigateway_models_2" {
  value = provider::arn::apigateway_models_2("rest-api-id")
}
