# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:policy-engine/policy-engine-id
output "bedrock_agentcore_policy_engine" {
  value = provider::arn::bedrock_agentcore_policy_engine("policy-engine-id")
}
