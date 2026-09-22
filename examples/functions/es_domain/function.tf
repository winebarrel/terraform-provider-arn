# arn:aws:es:ap-northeast-1:111111111111:domain/domain-name
output "es_domain" {
  value = provider::arn::es_domain("domain-name")
}
