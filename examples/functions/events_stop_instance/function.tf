# arn:aws:events:ap-northeast-1:111111111111:target/stop-instance
output "events_stop_instance" {
  value = provider::arn::events_stop_instance()
}
