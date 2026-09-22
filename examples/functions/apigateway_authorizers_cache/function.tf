# arn:aws:apigateway:ap-northeast-1::/apis/api-id/stages/stage-name/cache/authorizers
output "apigateway_authorizers_cache" {
  value = provider::arn::apigateway_authorizers_cache("api-id", "stage-name")
}
