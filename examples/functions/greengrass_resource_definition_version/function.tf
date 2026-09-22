# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/resources/resource-definition-id/versions/version-id
output "greengrass_resource_definition_version" {
  value = provider::arn::greengrass_resource_definition_version("resource-definition-id", "version-id")
}
