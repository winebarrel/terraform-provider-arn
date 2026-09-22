# arn:aws:sagemaker:ap-northeast-1:111111111111:image/image-name
output "sagemaker_image" {
  value = provider::arn::sagemaker_image("image-name")
}
