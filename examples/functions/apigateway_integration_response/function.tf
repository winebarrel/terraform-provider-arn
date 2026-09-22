# arn:aws:apigateway:ap-northeast-1::/apis/api-id/integrations/integration-id/integrationresponses/integration-response-id
output "apigateway_integration_response" {
  value = provider::arn::apigateway_integration_response("api-id", "integration-id", "integration-response-id")
}
