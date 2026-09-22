# arn:aws:b2bi:ap-northeast-1:111111111111:profile/resource-id
output "b2bi_profile" {
  value = provider::arn::b2bi_profile("resource-id")
}
