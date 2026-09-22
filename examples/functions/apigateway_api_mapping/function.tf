# arn:aws:apigateway:ap-northeast-1::/domainnames/domain-name/apimappings/api-mapping-id
output "apigateway_api_mapping" {
  value = provider::arn::apigateway_api_mapping("domain-name", "api-mapping-id")
}
