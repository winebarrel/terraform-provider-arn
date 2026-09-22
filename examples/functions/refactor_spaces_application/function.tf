# arn:aws:refactor-spaces:ap-northeast-1:111111111111:environment/environment-id/application/application-id
output "refactor_spaces_application" {
  value = provider::arn::refactor_spaces_application("environment-id", "application-id")
}
