# arn:aws:execute-api:ap-northeast-1:111111111111:/domainnames/domain-name+domain-identifier
output "execute_api_execute_api_domain" {
  value = provider::arn::execute_api_execute_api_domain("domain-name", "domain-identifier")
}
