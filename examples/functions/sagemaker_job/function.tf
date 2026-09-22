# arn:aws:sagemaker:ap-northeast-1:111111111111:job/job-category/job-name
output "sagemaker_job" {
  value = provider::arn::sagemaker_job("job-category", "job-name")
}
