# arn:aws:sagemaker:ap-northeast-1:111111111111:model-bias-job-definition/model-bias-job-definition-name
output "sagemaker_model_bias_job_definition" {
  value = provider::arn::sagemaker_model_bias_job_definition("model-bias-job-definition-name")
}
