# arn:aws:apigateway:ap-northeast-1::/apis/api-id/models
output "apigateway_models" {
  value = provider::arn::apigateway_models("api-id")
}
