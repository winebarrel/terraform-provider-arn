# arn:aws:b2bi:ap-northeast-1:111111111111:partnership/resource-id
output "b2bi_partnership" {
  value = provider::arn::b2bi_partnership("resource-id")
}
