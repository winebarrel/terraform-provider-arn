# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:consent-portal/consent-portal-id
output "bedrock_agentcore_consent_portal" {
  value = provider::arn::bedrock_agentcore_consent_portal("consent-portal-id")
}
