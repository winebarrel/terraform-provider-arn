# arn:aws:sqs:ap-northeast-1:111111111111:queue-name
output "sqs_queue" {
  value = provider::arn::sqs_queue("queue-name")
}
