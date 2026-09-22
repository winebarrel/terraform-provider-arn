# arn:aws:personalize:::feature-transformation/resource-id
output "personalize_feature_transformation" {
  value = provider::arn::personalize_feature_transformation("resource-id")
}
