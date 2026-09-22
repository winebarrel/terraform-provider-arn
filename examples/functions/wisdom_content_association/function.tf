# arn:aws:wisdom:ap-northeast-1:111111111111:content-association/knowledge-base-id/content-id/content-association-id
output "wisdom_content_association" {
  value = provider::arn::wisdom_content_association("knowledge-base-id", "content-id", "content-association-id")
}
