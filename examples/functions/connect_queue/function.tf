# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/queue/queue-id
output "connect_queue" {
  value = provider::arn::connect_queue("instance-id", "queue-id")
}
