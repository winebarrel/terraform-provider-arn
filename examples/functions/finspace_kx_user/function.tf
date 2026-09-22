# arn:aws:finspace:ap-northeast-1:111111111111:kxEnvironment/environment-id/kxUser/user-name
output "finspace_kx_user" {
  value = provider::arn::finspace_kx_user("environment-id", "user-name")
}
