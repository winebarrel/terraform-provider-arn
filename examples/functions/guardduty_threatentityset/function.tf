# arn:aws:guardduty:ap-northeast-1:111111111111:detector/detector-id/threatentityset/threat-entity-set-id
output "guardduty_threatentityset" {
  value = provider::arn::guardduty_threatentityset("detector-id", "threat-entity-set-id")
}
