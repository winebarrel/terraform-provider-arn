# arn:aws:datasync:ap-northeast-1:111111111111:system/storage-system-id
output "datasync_storagesystem" {
  value = provider::arn::datasync_storagesystem("storage-system-id")
}
