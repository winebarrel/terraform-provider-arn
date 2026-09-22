# arn:aws:mediaconvert:ap-northeast-1:111111111111:queues/queue-name
output "mediaconvert_queue" {
  value = provider::arn::mediaconvert_queue("queue-name")
}
