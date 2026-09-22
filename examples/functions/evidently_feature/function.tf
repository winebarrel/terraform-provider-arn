# arn:aws:evidently:ap-northeast-1:111111111111:project/project-name/feature/feature-name
output "evidently_feature" {
  value = provider::arn::evidently_feature("project-name", "feature-name")
}
