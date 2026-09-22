# arn:aws:payment-cryptography:ap-northeast-1:111111111111:key/key-id
output "payment_cryptography_key" {
  value = provider::arn::payment_cryptography_key("key-id")
}
