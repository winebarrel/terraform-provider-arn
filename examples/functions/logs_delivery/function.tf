# arn:aws:logs:ap-northeast-1:111111111111:delivery:delivery-name
output "logs_delivery" {
  value = provider::arn::logs_delivery("delivery-name")
}
