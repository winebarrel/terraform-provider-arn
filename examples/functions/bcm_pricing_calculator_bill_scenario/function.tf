# arn:aws:bcm-pricing-calculator::111111111111:bill-scenario/bill-scenario-id
output "bcm_pricing_calculator_bill_scenario" {
  value = provider::arn::bcm_pricing_calculator_bill_scenario("bill-scenario-id")
}
