# arn:aws:bedrock-agentcore:ap-northeast-1:aws:browser/browser-id
output "bedrock_agentcore_browser" {
  value = provider::arn::bedrock_agentcore_browser("browser-id")
}
