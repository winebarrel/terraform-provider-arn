# arn:aws:sagemaker:ap-northeast-1:111111111111:app-image-config/app-image-config-name
output "sagemaker_app_image_config" {
  value = provider::arn::sagemaker_app_image_config("app-image-config-name")
}
