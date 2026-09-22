# arn:aws:payment-cryptography:ap-northeast-1:111111111111:alias/alias
output "payment_cryptography_alias" {
  value = provider::arn::payment_cryptography_alias("alias")
}
