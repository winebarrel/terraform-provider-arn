# arn:aws:sagemaker:ap-northeast-1:111111111111:optimization-job/optimization-job-name
output "sagemaker_optimization_job" {
  value = provider::arn::sagemaker_optimization_job("optimization-job-name")
}
