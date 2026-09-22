# arn:aws:apigateway:ap-northeast-1::/apis/api-id/integrations/integration-id
output "apigateway_integration" {
  value = provider::arn::apigateway_integration("api-id", "integration-id")
}
