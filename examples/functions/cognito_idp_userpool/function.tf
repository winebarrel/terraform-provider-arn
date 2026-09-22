# arn:aws:cognito-idp:ap-northeast-1:111111111111:userpool/user-pool-id
output "cognito_idp_userpool" {
  value = provider::arn::cognito_idp_userpool("user-pool-id")
}
