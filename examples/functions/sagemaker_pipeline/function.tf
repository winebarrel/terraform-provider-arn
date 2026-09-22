# arn:aws:sagemaker:ap-northeast-1:111111111111:pipeline/pipeline-name
output "sagemaker_pipeline" {
  value = provider::arn::sagemaker_pipeline("pipeline-name")
}
