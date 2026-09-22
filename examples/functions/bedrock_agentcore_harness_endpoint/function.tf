# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:harness/harness-id/harness-endpoint/name
output "bedrock_agentcore_harness_endpoint" {
  value = provider::arn::bedrock_agentcore_harness_endpoint("harness-id", "name")
}
