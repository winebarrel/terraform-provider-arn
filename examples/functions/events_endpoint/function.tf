# arn:aws:events:ap-northeast-1:111111111111:endpoint/endpoint-name
output "events_endpoint" {
  value = provider::arn::events_endpoint("endpoint-name")
}
