# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:registry/registry-id/record/record-id
output "bedrock_agentcore_registry_record" {
  value = provider::arn::bedrock_agentcore_registry_record("registry-id", "record-id")
}
