# arn:aws:apigateway:ap-northeast-1::/apis/api-id/authorizers
output "apigateway_authorizers" {
  value = provider::arn::apigateway_authorizers("api-id")
}
