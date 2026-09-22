# arn:aws:iq:ap-northeast-1::paymentRequest/conversation-id/proposal-id/payment-request-id
output "iq_payment_request" {
  value = provider::arn::iq_payment_request("conversation-id", "proposal-id", "payment-request-id")
}
