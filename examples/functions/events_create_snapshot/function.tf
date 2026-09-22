# arn:aws:events:ap-northeast-1:111111111111:target/create-snapshot
output "events_create_snapshot" {
  value = provider::arn::events_create_snapshot()
}
