# arn:aws:logs:ap-northeast-1:111111111111:destination:destination-name
output "logs_destination" {
  value = provider::arn::logs_destination("destination-name")
}
