# arn:aws:sagemaker:ap-northeast-1:111111111111:transform-job/transform-job-name
output "sagemaker_transform_job" {
  value = provider::arn::sagemaker_transform_job("transform-job-name")
}
