# arn:aws:payments::111111111111:payment-instrument:resource-id
output "payments_payment_instrument" {
  value = provider::arn::payments_payment_instrument("resource-id")
}
