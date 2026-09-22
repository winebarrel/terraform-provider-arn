# arn:aws:sagemaker:ap-northeast-1:111111111111:model-package-group/model-package-group-name
output "sagemaker_model_package_group" {
  value = provider::arn::sagemaker_model_package_group("model-package-group-name")
}
