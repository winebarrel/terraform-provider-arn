# arn:aws:bugbust:ap-northeast-1:111111111111:events/event-id
output "bugbust_event" {
  value = provider::arn::bugbust_event("event-id")
}
