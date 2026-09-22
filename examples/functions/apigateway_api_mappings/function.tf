# arn:aws:apigateway:ap-northeast-1::/domainnames/domain-name/apimappings
output "apigateway_api_mappings" {
  value = provider::arn::apigateway_api_mappings("domain-name")
}
