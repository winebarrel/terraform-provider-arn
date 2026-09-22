# arn:aws:appstream:ap-northeast-1:111111111111:image-builder/image-builder-name
output "appstream_image_builder" {
  value = provider::arn::appstream_image_builder("image-builder-name")
}
