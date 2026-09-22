# arn:aws:apigateway:ap-northeast-1::/apis/api-id/cors
output "apigateway_cors" {
  value = provider::arn::apigateway_cors("api-id")
}
