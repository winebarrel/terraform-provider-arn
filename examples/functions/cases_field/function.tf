# arn:aws:cases:ap-northeast-1:111111111111:domain/domain-id/field/field-id
output "cases_field" {
  value = provider::arn::cases_field("domain-id", "field-id")
}
