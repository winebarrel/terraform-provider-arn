# arn:aws:batch:ap-northeast-1:111111111111:job-queue/job-queue-name/quota-share/quota-share-name
output "batch_quota_share" {
  value = provider::arn::batch_quota_share("job-queue-name", "quota-share-name")
}
