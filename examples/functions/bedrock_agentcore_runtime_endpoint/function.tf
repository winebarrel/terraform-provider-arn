# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:runtime/runtime-id/runtime-endpoint/name
output "bedrock_agentcore_runtime_endpoint" {
  value = provider::arn::bedrock_agentcore_runtime_endpoint("runtime-id", "name")
}
