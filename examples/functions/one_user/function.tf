# arn:aws:one:ap-northeast-1:111111111111:user/user-id
output "one_user" {
  value = provider::arn::one_user("user-id")
}
