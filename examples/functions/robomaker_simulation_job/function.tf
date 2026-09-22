# arn:aws:robomaker:ap-northeast-1:111111111111:simulation-job/simulation-job-id
output "robomaker_simulation_job" {
  value = provider::arn::robomaker_simulation_job("simulation-job-id")
}
