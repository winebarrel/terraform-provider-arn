# arn:aws:personalize:::recipe/resource-id
output "personalize_recipe" {
  value = provider::arn::personalize_recipe("resource-id")
}
