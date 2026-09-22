# arn:aws:wisdom:ap-northeast-1:111111111111:ai-agent/assistant-id/ai-agent-id
output "wisdom_ai_agent" {
  value = provider::arn::wisdom_ai_agent("assistant-id", "ai-agent-id")
}
