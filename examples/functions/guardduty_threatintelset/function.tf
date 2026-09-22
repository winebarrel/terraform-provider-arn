# arn:aws:guardduty:ap-northeast-1:111111111111:detector/detector-id/threatintelset/threat-intel-set-id
output "guardduty_threatintelset" {
  value = provider::arn::guardduty_threatintelset("detector-id", "threat-intel-set-id")
}
