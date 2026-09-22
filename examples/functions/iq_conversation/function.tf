# arn:aws:iq:ap-northeast-1::conversation/conversation-id
output "iq_conversation" {
  value = provider::arn::iq_conversation("conversation-id")
}
