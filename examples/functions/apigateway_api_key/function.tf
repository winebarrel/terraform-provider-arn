# arn:aws:apigateway:ap-northeast-1::/apikeys/api-key-id
output "apigateway_api_key" {
  value = provider::arn::apigateway_api_key("api-key-id")
}
