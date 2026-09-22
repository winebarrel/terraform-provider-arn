# arn:aws:a4b:ap-northeast-1:111111111111:gateway/resource-id
output "a4b_gateway" {
  value = provider::arn::a4b_gateway("resource-id")
}
