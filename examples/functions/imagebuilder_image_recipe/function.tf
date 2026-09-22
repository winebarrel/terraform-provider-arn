# arn:aws:imagebuilder:ap-northeast-1:111111111111:image-recipe/image-recipe-name/image-recipe-version
output "imagebuilder_image_recipe" {
  value = provider::arn::imagebuilder_image_recipe("image-recipe-name", "image-recipe-version")
}
