# arn:aws:apigateway:ap-northeast-1::/apis
output "apigateway_apis" {
  value = provider::arn::apigateway_apis()
}
