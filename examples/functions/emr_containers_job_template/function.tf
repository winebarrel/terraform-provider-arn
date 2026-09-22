# arn:aws:emr-containers:ap-northeast-1:111111111111:/jobtemplates/job-template-id
output "emr_containers_job_template" {
  value = provider::arn::emr_containers_job_template("job-template-id")
}
