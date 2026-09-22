# arn:aws:billing::111111111111:billingview/resource-id
output "ce_billingview" {
  value = provider::arn::ce_billingview("resource-id")
}
