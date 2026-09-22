# arn:aws:ses:ap-northeast-1:111111111111:addon-subscription/addon-subscription-id
output "ses_addon_subscription" {
  value = provider::arn::ses_addon_subscription("addon-subscription-id")
}
