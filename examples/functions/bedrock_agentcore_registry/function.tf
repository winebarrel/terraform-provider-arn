# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:registry/registry-id
output "bedrock_agentcore_registry" {
  value = provider::arn::bedrock_agentcore_registry("registry-id")
}
