# arn:aws:sqs:ap-northeast-1:111111111111:my-queue
output "queue" {
  value = provider::arn::sqs_queue("my-queue")
}

# arn:aws:sqs:us-east-1:111111111111:my-queue
output "queue_in_us" {
  value = provider::arn::sqs_queue("my-queue", { region = "us-east-1" })
}
