# arn:aws:billing::111111111111:billingview/resource-id
output "bcm_data_exports_billingview" {
  value = provider::arn::bcm_data_exports_billingview("resource-id")
}
