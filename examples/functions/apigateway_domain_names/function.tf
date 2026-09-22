# arn:aws:apigateway:ap-northeast-1::/domainnames
output "apigateway_domain_names" {
  value = provider::arn::apigateway_domain_names()
}
