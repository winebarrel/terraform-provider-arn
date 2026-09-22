# arn:aws:identity-sync:ap-northeast-1:111111111111:target/sync-profile-name/sync-target-name
output "identity_sync_sync_target_resource" {
  value = provider::arn::identity_sync_sync_target_resource("sync-profile-name", "sync-target-name")
}
