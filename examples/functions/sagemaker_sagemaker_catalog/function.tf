# arn:aws:sagemaker:ap-northeast-1:111111111111:sagemaker-catalog/resource-catalog-name
output "sagemaker_sagemaker_catalog" {
  value = provider::arn::sagemaker_sagemaker_catalog("resource-catalog-name")
}
