# arn:aws:freertos:ap-northeast-1:111111111111:subscription/subscription-id
output "freertos_subscription" {
  value = provider::arn::freertos_subscription("subscription-id")
}
