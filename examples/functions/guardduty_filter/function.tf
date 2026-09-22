# arn:aws:guardduty:ap-northeast-1:111111111111:detector/detector-id/filter/filter-name
output "guardduty_filter" {
  value = provider::arn::guardduty_filter("detector-id", "filter-name")
}
