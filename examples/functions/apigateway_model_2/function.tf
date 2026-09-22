# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/models/model-name
output "apigateway_model_2" {
  value = provider::arn::apigateway_model_2("rest-api-id", "model-name")
}
