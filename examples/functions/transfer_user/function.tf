# arn:aws:transfer:ap-northeast-1:111111111111:user/server-id/user-name
output "transfer_user" {
  value = provider::arn::transfer_user("server-id", "user-name")
}
