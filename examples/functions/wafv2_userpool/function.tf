# arn:aws:cognito-idp:ap-northeast-1:111111111111:userpool/user-pool-id
output "wafv2_userpool" {
  value = provider::arn::wafv2_userpool("user-pool-id")
}
