# arn:aws:dms:ap-northeast-1:111111111111:es:*
output "dms_event_subscription" {
  value = provider::arn::dms_event_subscription()
}
