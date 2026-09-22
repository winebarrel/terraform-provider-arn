# arn:aws:cognito-sync:ap-northeast-1:111111111111:identitypool/identity-pool-id/identity/identity-id
output "cognito_sync_identity" {
  value = provider::arn::cognito_sync_identity("identity-pool-id", "identity-id")
}
