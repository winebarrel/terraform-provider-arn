# arn:aws:b2bi:ap-northeast-1:111111111111:transformer/resource-id
output "b2bi_transformer" {
  value = provider::arn::b2bi_transformer("resource-id")
}
