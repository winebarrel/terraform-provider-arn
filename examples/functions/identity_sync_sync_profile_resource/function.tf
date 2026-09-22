# arn:aws:identity-sync:ap-northeast-1:111111111111:profile/sync-profile-name
output "identity_sync_sync_profile_resource" {
  value = provider::arn::identity_sync_sync_profile_resource("sync-profile-name")
}
