# arn:aws:finspace:ap-northeast-1:111111111111:user/user-id
output "finspace_user" {
  value = provider::arn::finspace_user("user-id")
}
