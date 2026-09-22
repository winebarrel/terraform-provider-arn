# arn:aws:sagemaker:ap-northeast-1:111111111111:automl-job/auto-ml-job-job-name
output "sagemaker_automl_job" {
  value = provider::arn::sagemaker_automl_job("auto-ml-job-job-name")
}
