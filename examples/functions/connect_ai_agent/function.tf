# arn:aws:wisdom:ap-northeast-1:111111111111:ai-agent/assistant-id/ai-agent-id:version
output "connect_ai_agent" {
  value = provider::arn::connect_ai_agent("assistant-id", "ai-agent-id", "version")
}
