# arn:aws:sagemaker:ap-northeast-1:111111111111:model-package/model-package-name
output "sagemaker_model_package" {
  value = provider::arn::sagemaker_model_package("model-package-name")
}
