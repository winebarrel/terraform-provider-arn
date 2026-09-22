# arn:aws:refactor-spaces:ap-northeast-1:111111111111:environment/environment-id/application/application-id/route/route-id
output "refactor_spaces_route" {
  value = provider::arn::refactor_spaces_route("environment-id", "application-id", "route-id")
}
