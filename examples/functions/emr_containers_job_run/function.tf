# arn:aws:emr-containers:ap-northeast-1:111111111111:/virtualclusters/virtual-cluster-id/jobruns/job-run-id
output "emr_containers_job_run" {
  value = provider::arn::emr_containers_job_run("virtual-cluster-id", "job-run-id")
}
