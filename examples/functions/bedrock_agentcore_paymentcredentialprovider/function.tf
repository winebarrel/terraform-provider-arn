# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:token-vault/token-vault-id/paymentcredentialprovider/name
output "bedrock_agentcore_paymentcredentialprovider" {
  value = provider::arn::bedrock_agentcore_paymentcredentialprovider("token-vault-id", "name")
}
