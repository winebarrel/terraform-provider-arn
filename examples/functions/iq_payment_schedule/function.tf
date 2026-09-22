# arn:aws:iq:ap-northeast-1::paymentSchedule/conversation-id/proposal-id/version-id
output "iq_payment_schedule" {
  value = provider::arn::iq_payment_schedule("conversation-id", "proposal-id", "version-id")
}
