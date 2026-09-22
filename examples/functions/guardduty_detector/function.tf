# arn:aws:guardduty:ap-northeast-1:111111111111:detector/detector-id
output "guardduty_detector" {
  value = provider::arn::guardduty_detector("detector-id")
}
