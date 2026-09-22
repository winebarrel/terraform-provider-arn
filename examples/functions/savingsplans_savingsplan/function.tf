# arn:aws:savingsplans::111111111111:savingsplan/resource-id
output "savingsplans_savingsplan" {
  value = provider::arn::savingsplans_savingsplan("resource-id")
}
