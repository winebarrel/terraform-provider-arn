# arn:aws:wisdom:ap-northeast-1:111111111111:knowledge-base/knowledge-base-id
output "wisdom_knowledge_base" {
  value = provider::arn::wisdom_knowledge_base("knowledge-base-id")
}
