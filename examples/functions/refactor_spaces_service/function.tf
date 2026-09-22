# arn:aws:refactor-spaces:ap-northeast-1:111111111111:environment/environment-id/application/application-id/service/service-id
output "refactor_spaces_service" {
  value = provider::arn::refactor_spaces_service("environment-id", "application-id", "service-id")
}
