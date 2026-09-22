# arn:aws:imagebuilder:ap-northeast-1:111111111111:image/image-name/image-version/*
output "imagebuilder_all_image_build_versions" {
  value = provider::arn::imagebuilder_all_image_build_versions("image-name", "image-version")
}
