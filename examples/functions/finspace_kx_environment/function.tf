# arn:aws:finspace:ap-northeast-1:111111111111:kxEnvironment/environment-id
output "finspace_kx_environment" {
  value = provider::arn::finspace_kx_environment("environment-id")
}
