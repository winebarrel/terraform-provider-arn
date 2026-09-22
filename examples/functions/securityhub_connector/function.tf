# arn:aws:securityhub:ap-northeast-1:111111111111:connector/connector-id
output "securityhub_connector" {
  value = provider::arn::securityhub_connector("connector-id")
}
