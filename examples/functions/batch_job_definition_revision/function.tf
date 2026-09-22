# arn:aws:batch:ap-northeast-1:111111111111:job-definition/job-definition-name:revision
output "batch_job_definition_revision" {
  value = provider::arn::batch_job_definition_revision("job-definition-name", "revision")
}
