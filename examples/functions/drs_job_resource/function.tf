# arn:aws:drs:ap-northeast-1:111111111111:job/job-id
output "drs_job_resource" {
  value = provider::arn::drs_job_resource("job-id")
}
