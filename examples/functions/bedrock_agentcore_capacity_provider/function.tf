# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:capacity-provider/capacity-provider-id
output "bedrock_agentcore_capacity_provider" {
  value = provider::arn::bedrock_agentcore_capacity_provider("capacity-provider-id")
}
