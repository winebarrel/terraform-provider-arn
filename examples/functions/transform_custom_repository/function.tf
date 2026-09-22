# arn:aws:transform-custom:ap-northeast-1:111111111111:repository/repository-id
output "transform_custom_repository" {
  value = provider::arn::transform_custom_repository("repository-id")
}
