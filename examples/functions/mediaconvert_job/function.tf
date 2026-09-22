# arn:aws:mediaconvert:ap-northeast-1:111111111111:jobs/job-id
output "mediaconvert_job" {
  value = provider::arn::mediaconvert_job("job-id")
}
