# arn:aws:deadline:ap-northeast-1:111111111111:farm/farm-id/queue/queue-id/job/job-id
output "deadline_job" {
  value = provider::arn::deadline_job("farm-id", "queue-id", "job-id")
}
