# arn:aws:guardduty:ap-northeast-1:111111111111:detector/detector-id/trustedentityset/trusted-entity-set-id
output "guardduty_trustedentityset" {
  value = provider::arn::guardduty_trustedentityset("detector-id", "trusted-entity-set-id")
}
