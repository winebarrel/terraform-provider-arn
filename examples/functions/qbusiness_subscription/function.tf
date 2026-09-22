# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/subscription/subscription-id
output "qbusiness_subscription" {
  value = provider::arn::qbusiness_subscription("application-id", "subscription-id")
}
