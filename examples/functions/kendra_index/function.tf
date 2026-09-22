# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id
output "kendra_index" {
  value = provider::arn::kendra_index("index-id")
}
