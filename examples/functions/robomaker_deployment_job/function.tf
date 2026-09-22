# arn:aws:robomaker:ap-northeast-1:111111111111:deployment-job/deployment-job-id
output "robomaker_deployment_job" {
  value = provider::arn::robomaker_deployment_job("deployment-job-id")
}
