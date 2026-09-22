# arn:aws:sagemaker:ap-northeast-1:111111111111:data-quality-job-definition/data-quality-job-definition-name
output "sagemaker_data_quality_job_definition" {
  value = provider::arn::sagemaker_data_quality_job_definition("data-quality-job-definition-name")
}
