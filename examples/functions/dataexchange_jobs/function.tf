# arn:aws:dataexchange:ap-northeast-1:111111111111:jobs/job-id
output "dataexchange_jobs" {
  value = provider::arn::dataexchange_jobs("job-id")
}
