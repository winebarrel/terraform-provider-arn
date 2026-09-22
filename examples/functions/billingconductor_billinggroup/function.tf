# arn:aws:billingconductor::111111111111:billinggroup/billing-group-id
output "billingconductor_billinggroup" {
  value = provider::arn::billingconductor_billinggroup("billing-group-id")
}
