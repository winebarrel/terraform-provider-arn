# arn:aws:mgn:ap-northeast-1:111111111111:connector/connector-id
output "mgn_connector_resource" {
  value = provider::arn::mgn_connector_resource("connector-id")
}
