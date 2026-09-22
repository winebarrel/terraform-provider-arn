# arn:aws:mediatailor:ap-northeast-1:111111111111:prefetchSchedule/resource-id
output "mediatailor_prefetch_schedule" {
  value = provider::arn::mediatailor_prefetch_schedule("resource-id")
}
