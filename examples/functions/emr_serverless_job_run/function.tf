# arn:aws:emr-serverless:ap-northeast-1:111111111111:/applications/application-id/jobruns/job-run-id
output "emr_serverless_job_run" {
  value = provider::arn::emr_serverless_job_run("application-id", "job-run-id")
}
