# arn:aws:apigateway:ap-northeast-1::/domainnames/domain-name
output "apigateway_domain_name" {
  value = provider::arn::apigateway_domain_name("domain-name")
}
