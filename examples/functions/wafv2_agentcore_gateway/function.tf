# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:gateway/gateway-id
output "wafv2_agentcore_gateway" {
  value = provider::arn::wafv2_agentcore_gateway("gateway-id")
}
