# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:runtime/runtime-id
output "bedrock_agentcore_runtime" {
  value = provider::arn::bedrock_agentcore_runtime("runtime-id")
}
