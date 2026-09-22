# arn:aws:apigateway:ap-northeast-1:111111111111:/domainnames/domain-name+domain-identifier
output "apigateway_private_domain_name" {
  value = provider::arn::apigateway_private_domain_name("domain-name", "domain-identifier")
}
