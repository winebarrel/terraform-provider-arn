# arn:aws:apigateway:ap-northeast-1::/apis/api-id/stages
output "apigateway_stages" {
  value = provider::arn::apigateway_stages("api-id")
}
