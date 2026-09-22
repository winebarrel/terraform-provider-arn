# arn:aws:s3:ap-northeast-1:111111111111:job/job-id
output "s3_job" {
  value = provider::arn::s3_job("job-id")
}
