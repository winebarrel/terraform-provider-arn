# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/connectors/connector-definition-id/versions/version-id
output "greengrass_connector_definition_version" {
  value = provider::arn::greengrass_connector_definition_version("connector-definition-id", "version-id")
}
