# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:browser-custom/browser-id
output "bedrock_agentcore_browser_custom" {
  value = provider::arn::bedrock_agentcore_browser_custom("browser-id")
}
