# arn:aws:databrew:ap-northeast-1:111111111111:recipe/resource-id
output "databrew_recipe" {
  value = provider::arn::databrew_recipe("resource-id")
}
