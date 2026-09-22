# arn:aws:sagemaker:ap-northeast-1:111111111111:pipeline/pipeline-name/execution/random-string
output "sagemaker_pipeline_execution" {
  value = provider::arn::sagemaker_pipeline_execution("pipeline-name", "random-string")
}
