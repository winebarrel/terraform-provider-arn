# arn:aws:iottwinmaker:ap-northeast-1:111111111111:workspace/workspace-id/sync-job/sync-job-id
output "iottwinmaker_sync_job" {
  value = provider::arn::iottwinmaker_sync_job("workspace-id", "sync-job-id")
}
