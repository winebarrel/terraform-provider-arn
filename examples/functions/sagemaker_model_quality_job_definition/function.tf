# arn:aws:sagemaker:ap-northeast-1:111111111111:model-quality-job-definition/model-quality-job-definition-name
output "sagemaker_model_quality_job_definition" {
  value = provider::arn::sagemaker_model_quality_job_definition("model-quality-job-definition-name")
}
