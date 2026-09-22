# arn:aws:amplify:ap-northeast-1:111111111111:apps/app-id
output "wafv2_amplify_app" {
  value = provider::arn::wafv2_amplify_app("app-id")
}
