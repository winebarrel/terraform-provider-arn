# arn:aws:memorydb:ap-northeast-1:111111111111:user/user-name
output "memorydb_user" {
  value = provider::arn::memorydb_user("user-name")
}
