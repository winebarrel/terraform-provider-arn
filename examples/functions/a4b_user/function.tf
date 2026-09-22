# arn:aws:a4b:ap-northeast-1:111111111111:user/resource-id
output "a4b_user" {
  value = provider::arn::a4b_user("resource-id")
}
