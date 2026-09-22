# arn:aws:cases:ap-northeast-1:111111111111:domain/domain-id
output "cases_domain" {
  value = provider::arn::cases_domain("domain-id")
}
