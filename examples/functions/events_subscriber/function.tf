# arn:aws:events:ap-northeast-1:111111111111:subscriber/subscriber-name/opaque-id
output "events_subscriber" {
  value = provider::arn::events_subscriber("subscriber-name", "opaque-id")
}
