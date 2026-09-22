# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:token-vault/token-vault-id
output "bedrock_agentcore_token_vault" {
  value = provider::arn::bedrock_agentcore_token_vault("token-vault-id")
}
