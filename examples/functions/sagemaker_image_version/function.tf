# arn:aws:sagemaker:ap-northeast-1:111111111111:image-version/image-name/version
output "sagemaker_image_version" {
  value = provider::arn::sagemaker_image_version("image-name", "version")
}
