# arn:aws:finspace:ap-northeast-1:111111111111:environment/environment-id
output "finspace_environment" {
  value = provider::arn::finspace_environment("environment-id")
}
