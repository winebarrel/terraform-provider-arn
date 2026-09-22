# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:configuration-bundle/configuration-bundle-id
output "bedrock_agentcore_configuration_bundle" {
  value = provider::arn::bedrock_agentcore_configuration_bundle("configuration-bundle-id")
}
