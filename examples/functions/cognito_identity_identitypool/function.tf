# arn:aws:cognito-identity:ap-northeast-1:111111111111:identitypool/identity-pool-id
output "cognito_identity_identitypool" {
  value = provider::arn::cognito_identity_identitypool("identity-pool-id")
}
