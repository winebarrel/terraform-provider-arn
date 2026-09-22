# arn:aws:transfer:ap-northeast-1:111111111111:connector/connector-id
output "transfer_connector" {
  value = provider::arn::transfer_connector("connector-id")
}
