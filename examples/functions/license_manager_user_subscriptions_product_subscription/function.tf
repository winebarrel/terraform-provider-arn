# arn:aws:license-manager-user-subscriptions:ap-northeast-1:111111111111:product-subscription/product-subscription-id
output "license_manager_user_subscriptions_product_subscription" {
  value = provider::arn::license_manager_user_subscriptions_product_subscription("product-subscription-id")
}
