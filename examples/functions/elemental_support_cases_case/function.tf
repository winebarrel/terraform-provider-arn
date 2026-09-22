# arn:aws:elemental-support-cases::111111111111:case/resource-id
output "elemental_support_cases_case" {
  value = provider::arn::elemental_support_cases_case("resource-id")
}
