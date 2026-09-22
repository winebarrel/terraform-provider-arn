# arn:aws:refactor-spaces:ap-northeast-1:111111111111:environment/environment-id
output "refactor_spaces_environment" {
  value = provider::arn::refactor_spaces_environment("environment-id")
}
