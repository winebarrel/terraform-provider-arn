# arn:aws:appstream:ap-northeast-1:111111111111:image/image-name
output "appstream_image" {
  value = provider::arn::appstream_image("image-name")
}
