# arn:aws:bedrock:ap-northeast-1:111111111111:knowledge-base/knowledge-base-id
output "bedrock_knowledge_base" {
  value = provider::arn::bedrock_knowledge_base("knowledge-base-id")
}
