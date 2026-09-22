# arn:aws:cases:ap-northeast-1:111111111111:domain/domain-id/case-rule/case-rule-id
output "cases_case_rule" {
  value = provider::arn::cases_case_rule("domain-id", "case-rule-id")
}
