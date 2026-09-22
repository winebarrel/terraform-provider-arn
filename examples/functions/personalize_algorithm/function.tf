# arn:aws:personalize:::algorithm/resource-id
output "personalize_algorithm" {
  value = provider::arn::personalize_algorithm("resource-id")
}
