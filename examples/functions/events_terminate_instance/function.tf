# arn:aws:events:ap-northeast-1:111111111111:target/terminate-instance
output "events_terminate_instance" {
  value = provider::arn::events_terminate_instance()
}
