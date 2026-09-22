# arn:aws:cognito-sync:ap-northeast-1:111111111111:identitypool/identity-pool-id
output "cognito_sync_identitypool" {
  value = provider::arn::cognito_sync_identitypool("identity-pool-id")
}
