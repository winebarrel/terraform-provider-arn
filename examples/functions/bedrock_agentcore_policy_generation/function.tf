# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:policy-engine/policy-engine-id/policy-generation/policy-generation-id
output "bedrock_agentcore_policy_generation" {
  value = provider::arn::bedrock_agentcore_policy_generation("policy-engine-id", "policy-generation-id")
}
