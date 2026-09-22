# arn:aws:apigateway:ap-northeast-1::/apis/api-id/integrations
output "apigateway_integrations" {
  value = provider::arn::apigateway_integrations("api-id")
}
