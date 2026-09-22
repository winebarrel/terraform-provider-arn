# arn:aws:personalize:ap-northeast-1:111111111111:batch-segment-job/resource-id
output "personalize_batch_segment_job" {
  value = provider::arn::personalize_batch_segment_job("resource-id")
}
