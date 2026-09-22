# arn:aws:glue:ap-northeast-1:111111111111:job/job-name
output "glue_job" {
  value = provider::arn::glue_job("job-name")
}
