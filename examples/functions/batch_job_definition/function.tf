# arn:aws:batch:ap-northeast-1:111111111111:job-definition/job-definition-name
output "batch_job_definition" {
  value = provider::arn::batch_job_definition("job-definition-name")
}
