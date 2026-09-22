# arn:aws:imagebuilder:ap-northeast-1:111111111111:container-recipe/container-recipe-name/container-recipe-version
output "imagebuilder_container_recipe" {
  value = provider::arn::imagebuilder_container_recipe("container-recipe-name", "container-recipe-version")
}
