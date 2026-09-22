# arn:aws:a4b:ap-northeast-1:111111111111:profile/resource-id
output "a4b_profile" {
  value = provider::arn::a4b_profile("resource-id")
}
