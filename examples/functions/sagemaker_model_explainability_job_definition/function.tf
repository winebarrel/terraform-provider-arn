# arn:aws:sagemaker:ap-northeast-1:111111111111:model-explainability-job-definition/model-explainability-job-definition-name
output "sagemaker_model_explainability_job_definition" {
  value = provider::arn::sagemaker_model_explainability_job_definition("model-explainability-job-definition-name")
}
