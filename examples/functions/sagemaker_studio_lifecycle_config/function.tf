# arn:aws:sagemaker:ap-northeast-1:111111111111:studio-lifecycle-config/studio-lifecycle-config-name
output "sagemaker_studio_lifecycle_config" {
  value = provider::arn::sagemaker_studio_lifecycle_config("studio-lifecycle-config-name")
}
