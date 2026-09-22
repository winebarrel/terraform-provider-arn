# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:memory/memory-id
output "bedrock_agentcore_memory" {
  value = provider::arn::bedrock_agentcore_memory("memory-id")
}
