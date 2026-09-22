# arn:aws:glue:ap-northeast-1:111111111111:integrationresourceproperty/resource-type/resource-name
output "glue_integration_resource_property" {
  value = provider::arn::glue_integration_resource_property("resource-type", "resource-name")
}
