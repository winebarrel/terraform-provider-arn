# arn:aws:securityhub:ap-northeast-1:111111111111:connectorv2/connector-v2-id
output "securityhub_connectorv2" {
  value = provider::arn::securityhub_connectorv2("connector-v2-id")
}
