# arn:aws:billingconductor::111111111111:pricingrule/pricing-rule-id
output "billingconductor_pricingrule" {
  value = provider::arn::billingconductor_pricingrule("pricing-rule-id")
}
