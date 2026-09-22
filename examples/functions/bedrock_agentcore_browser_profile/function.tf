# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:browser-profile/browser-profile-id
output "bedrock_agentcore_browser_profile" {
  value = provider::arn::bedrock_agentcore_browser_profile("browser-profile-id")
}
