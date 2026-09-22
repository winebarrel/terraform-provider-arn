# arn:aws:elastictranscoder:ap-northeast-1:111111111111:job/job-id
output "elastictranscoder_job" {
  value = provider::arn::elastictranscoder_job("job-id")
}
