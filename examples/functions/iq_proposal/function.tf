# arn:aws:iq:ap-northeast-1::proposal/conversation-id/proposal-id
output "iq_proposal" {
  value = provider::arn::iq_proposal("conversation-id", "proposal-id")
}
