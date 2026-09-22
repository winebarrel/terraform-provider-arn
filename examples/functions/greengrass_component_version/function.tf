# arn:aws:greengrass:ap-northeast-1:111111111111:components:component-name:versions:component-version
output "greengrass_component_version" {
  value = provider::arn::greengrass_component_version("component-name", "component-version")
}
