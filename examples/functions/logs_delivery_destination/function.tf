# arn:aws:logs:ap-northeast-1:111111111111:delivery-destination:delivery-destination-name
output "logs_delivery_destination" {
  value = provider::arn::logs_delivery_destination("delivery-destination-name")
}
