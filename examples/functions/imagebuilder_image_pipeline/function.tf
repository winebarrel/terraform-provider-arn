# arn:aws:imagebuilder:ap-northeast-1:111111111111:image-pipeline/image-pipeline-name
output "imagebuilder_image_pipeline" {
  value = provider::arn::imagebuilder_image_pipeline("image-pipeline-name")
}
