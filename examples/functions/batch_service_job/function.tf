# arn:aws:batch:ap-northeast-1:111111111111:service-job/job-id
output "batch_service_job" {
  value = provider::arn::batch_service_job("job-id")
}
