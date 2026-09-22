# arn:aws:guardduty:ap-northeast-1:111111111111:detector/detector-id/ipset/ip-set-id
output "guardduty_ipset" {
  value = provider::arn::guardduty_ipset("detector-id", "ip-set-id")
}
