# arn:aws:wisdom:ap-northeast-1:111111111111:ai-prompt/assistant-id/ai-prompt-id
output "wisdom_ai_prompt" {
  value = provider::arn::wisdom_ai_prompt("assistant-id", "ai-prompt-id")
}
