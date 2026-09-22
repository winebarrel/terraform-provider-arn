# arn:aws:cognito-sync:ap-northeast-1:111111111111:identitypool/identity-pool-id/identity/identity-id/dataset/dataset-name
output "cognito_sync_dataset" {
  value = provider::arn::cognito_sync_dataset("identity-pool-id", "identity-id", "dataset-name")
}
