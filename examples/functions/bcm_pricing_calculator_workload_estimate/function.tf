# arn:aws:bcm-pricing-calculator::111111111111:workload-estimate/workload-estimate-id
output "bcm_pricing_calculator_workload_estimate" {
  value = provider::arn::bcm_pricing_calculator_workload_estimate("workload-estimate-id")
}
