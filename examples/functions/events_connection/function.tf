# arn:aws:events:ap-northeast-1:111111111111:connection/connection-name
output "events_connection" {
  value = provider::arn::events_connection("connection-name")
}
