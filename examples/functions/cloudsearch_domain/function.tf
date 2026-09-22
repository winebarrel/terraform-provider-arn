# arn:aws:cloudsearch:ap-northeast-1:111111111111:domain/domain-name
output "cloudsearch_domain" {
  value = provider::arn::cloudsearch_domain("domain-name")
}
