# arn:aws:apigateway:ap-northeast-1::/apis/api-id/integrations/integration-id/integrationresponses
output "apigateway_integration_responses" {
  value = provider::arn::apigateway_integration_responses("api-id", "integration-id")
}
