# arn:aws:sagemaker:ap-northeast-1:111111111111:context/context-name
output "sagemaker_context" {
  value = provider::arn::sagemaker_context("context-name")
}
