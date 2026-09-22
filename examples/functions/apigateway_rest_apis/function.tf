# arn:aws:apigateway:ap-northeast-1::/restapis
output "apigateway_rest_apis" {
  value = provider::arn::apigateway_rest_apis()
}
