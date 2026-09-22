# arn:aws:billingconductor::111111111111:pricingplan/pricing-plan-id
output "billingconductor_pricingplan" {
  value = provider::arn::billingconductor_pricingplan("pricing-plan-id")
}
