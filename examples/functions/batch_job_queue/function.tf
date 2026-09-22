# arn:aws:batch:ap-northeast-1:111111111111:job-queue/job-queue-name
output "batch_job_queue" {
  value = provider::arn::batch_job_queue("job-queue-name")
}
