# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:harness/harness-id
output "bedrock_agentcore_harness" {
  value = provider::arn::bedrock_agentcore_harness("harness-id")
}
