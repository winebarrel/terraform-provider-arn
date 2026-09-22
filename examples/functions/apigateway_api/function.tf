# arn:aws:apigateway:ap-northeast-1::/apis/api-id
output "apigateway_api" {
  value = provider::arn::apigateway_api("api-id")
}
