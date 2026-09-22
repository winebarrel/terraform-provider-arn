# arn:aws:pricingplanmanager::111111111111:subscription/subscription-id
output "pricingplanmanager_subscription" {
  value = provider::arn::pricingplanmanager_subscription("subscription-id")
}
