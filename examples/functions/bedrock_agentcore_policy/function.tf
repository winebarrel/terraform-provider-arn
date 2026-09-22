# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:policy-engine/policy-engine-id/policy/policy-id
output "bedrock_agentcore_policy" {
  value = provider::arn::bedrock_agentcore_policy("policy-engine-id", "policy-id")
}
