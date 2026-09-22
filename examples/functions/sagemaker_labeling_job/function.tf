# arn:aws:sagemaker:ap-northeast-1:111111111111:labeling-job/labeling-job-name
output "sagemaker_labeling_job" {
  value = provider::arn::sagemaker_labeling_job("labeling-job-name")
}
