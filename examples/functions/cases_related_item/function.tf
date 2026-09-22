# arn:aws:cases:ap-northeast-1:111111111111:domain/domain-id/case/case-id/related-item/related-item-id
output "cases_related_item" {
  value = provider::arn::cases_related_item("domain-id", "case-id", "related-item-id")
}
