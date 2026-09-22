# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:gateway/gateway-id
output "bedrock_agentcore_gateway" {
  value = provider::arn::bedrock_agentcore_gateway("gateway-id")
}
