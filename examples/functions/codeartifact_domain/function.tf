# arn:aws:codeartifact:ap-northeast-1:111111111111:domain/domain-name
output "codeartifact_domain" {
  value = provider::arn::codeartifact_domain("domain-name")
}
