# arn:aws:datasync:ap-northeast-1:111111111111:system/storage-system-id/job/discovery-job-id
output "datasync_discoveryjob" {
  value = provider::arn::datasync_discoveryjob("storage-system-id", "discovery-job-id")
}
