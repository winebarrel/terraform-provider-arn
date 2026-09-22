# arn:aws:wisdom:ap-northeast-1:111111111111:content/knowledge-base-id/content-id
output "wisdom_content" {
  value = provider::arn::wisdom_content("knowledge-base-id", "content-id")
}
