# arn:aws:deadline:ap-northeast-1:111111111111:farm/farm-id/queue/queue-id
output "deadline_queue" {
  value = provider::arn::deadline_queue("farm-id", "queue-id")
}
