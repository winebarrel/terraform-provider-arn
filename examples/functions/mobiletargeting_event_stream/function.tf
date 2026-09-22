# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/eventstream
output "mobiletargeting_event_stream" {
  value = provider::arn::mobiletargeting_event_stream("app-id")
}
