# arn:aws:wisdom:ap-northeast-1:111111111111:assistant/assistant-id
output "wisdom_assistant" {
  value = provider::arn::wisdom_assistant("assistant-id")
}
