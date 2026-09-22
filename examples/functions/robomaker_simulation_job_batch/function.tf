# arn:aws:robomaker:ap-northeast-1:111111111111:simulation-job-batch/simulation-job-batch-id
output "robomaker_simulation_job_batch" {
  value = provider::arn::robomaker_simulation_job_batch("simulation-job-batch-id")
}
