# arn:aws:mgn:ap-northeast-1:111111111111:job/job-id
output "mgn_job_resource" {
  value = provider::arn::mgn_job_resource("job-id")
}
