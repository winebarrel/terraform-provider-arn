# arn:aws:apigateway:ap-northeast-1::/apikeys
output "apigateway_api_keys" {
  value = provider::arn::apigateway_api_keys()
}
