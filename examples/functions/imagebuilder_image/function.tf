# arn:aws:imagebuilder:ap-northeast-1:111111111111:image/image-name/image-version/image-build-version
output "imagebuilder_image" {
  value = provider::arn::imagebuilder_image("image-name", "image-version", "image-build-version")
}
