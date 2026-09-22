# arn:aws:bcm-pricing-calculator::111111111111:bill-estimate/bill-estimate-id
output "bcm_pricing_calculator_bill_estimate" {
  value = provider::arn::bcm_pricing_calculator_bill_estimate("bill-estimate-id")
}
