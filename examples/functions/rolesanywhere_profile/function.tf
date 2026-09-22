# arn:aws:rolesanywhere:ap-northeast-1:111111111111:profile/profile-id
output "rolesanywhere_profile" {
  value = provider::arn::rolesanywhere_profile("profile-id")
}
