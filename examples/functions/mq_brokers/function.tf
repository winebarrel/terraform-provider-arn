# arn:aws:mq:ap-northeast-1:111111111111:broker:broker-name:broker-id
output "mq_brokers" {
  value = provider::arn::mq_brokers("broker-name", "broker-id")
}
