# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:token-vault/token-vault-id/oauth2credentialprovider/name
output "bedrock_agentcore_oauth2credentialprovider" {
  value = provider::arn::bedrock_agentcore_oauth2credentialprovider("token-vault-id", "name")
}
