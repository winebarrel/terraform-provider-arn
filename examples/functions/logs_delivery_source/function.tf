# arn:aws:logs:ap-northeast-1:111111111111:delivery-source:delivery-source-name
output "logs_delivery_source" {
  value = provider::arn::logs_delivery_source("delivery-source-name")
}
