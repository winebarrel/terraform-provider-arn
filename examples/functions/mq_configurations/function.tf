# arn:aws:mq:ap-northeast-1:111111111111:configuration:configuration-id
output "mq_configurations" {
  value = provider::arn::mq_configurations("configuration-id")
}
