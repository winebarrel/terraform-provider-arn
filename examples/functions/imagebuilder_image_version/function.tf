# arn:aws:imagebuilder:ap-northeast-1:111111111111:image/image-name/image-version
output "imagebuilder_image_version" {
  value = provider::arn::imagebuilder_image_version("image-name", "image-version")
}
