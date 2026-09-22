# arn:aws:frauddetector:ap-northeast-1:111111111111:event-type/resource-path
output "frauddetector_event_type" {
  value = provider::arn::frauddetector_event_type("resource-path")
}
