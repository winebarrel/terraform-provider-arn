# arn:aws:sagemaker:ap-northeast-1:111111111111:processing-job/processing-job-name
output "sagemaker_processing_job" {
  value = provider::arn::sagemaker_processing_job("processing-job-name")
}
