# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/resources/resource-definition-id
output "greengrass_resource_definition" {
  value = provider::arn::greengrass_resource_definition("resource-definition-id")
}
