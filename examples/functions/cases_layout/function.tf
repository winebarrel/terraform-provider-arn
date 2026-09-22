# arn:aws:cases:ap-northeast-1:111111111111:domain/domain-id/layout/layout-id
output "cases_layout" {
  value = provider::arn::cases_layout("domain-id", "layout-id")
}
