# arn:aws:b2bi:ap-northeast-1:111111111111:capability/resource-id
output "b2bi_capability" {
  value = provider::arn::b2bi_capability("resource-id")
}
