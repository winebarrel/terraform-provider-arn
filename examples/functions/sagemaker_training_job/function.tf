# arn:aws:sagemaker:ap-northeast-1:111111111111:training-job/training-job-name
output "sagemaker_training_job" {
  value = provider::arn::sagemaker_training_job("training-job-name")
}
