# arn:aws:wisdom:ap-northeast-1:111111111111:association/assistant-id/assistant-association-id
output "wisdom_assistant_association" {
  value = provider::arn::wisdom_assistant_association("assistant-id", "assistant-association-id")
}
