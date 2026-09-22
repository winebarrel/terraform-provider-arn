# arn:aws:lambda:ap-northeast-1:111111111111:network-connector:network-connector-id
output "lambda_network_connector" {
  value = provider::arn::lambda_network_connector("network-connector-id")
}
