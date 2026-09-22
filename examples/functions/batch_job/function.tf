# arn:aws:batch:ap-northeast-1:111111111111:job/job-id
output "batch_job" {
  value = provider::arn::batch_job("job-id")
}
