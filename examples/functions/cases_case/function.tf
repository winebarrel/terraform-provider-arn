# arn:aws:cases:ap-northeast-1:111111111111:domain/domain-id/case/case-id
output "cases_case" {
  value = provider::arn::cases_case("domain-id", "case-id")
}
