# arn:aws:billing::111111111111:billingview/resource-id
output "billing_billingview" {
  value = provider::arn::billing_billingview("resource-id")
}
