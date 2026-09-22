# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/connectors/connector-definition-id
output "greengrass_connector_definition" {
  value = provider::arn::greengrass_connector_definition("connector-definition-id")
}
