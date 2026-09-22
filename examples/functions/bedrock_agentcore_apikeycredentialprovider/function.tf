# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:token-vault/token-vault-id/apikeycredentialprovider/name
output "bedrock_agentcore_apikeycredentialprovider" {
  value = provider::arn::bedrock_agentcore_apikeycredentialprovider("token-vault-id", "name")
}
