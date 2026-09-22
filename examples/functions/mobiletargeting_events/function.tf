# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/events
output "mobiletargeting_events" {
  value = provider::arn::mobiletargeting_events("app-id")
}
