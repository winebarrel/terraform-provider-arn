# arn:aws:billingconductor::111111111111:customlineitem/custom-line-item-id
output "billingconductor_customlineitem" {
  value = provider::arn::billingconductor_customlineitem("custom-line-item-id")
}
