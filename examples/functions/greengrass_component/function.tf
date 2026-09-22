# arn:aws:greengrass:ap-northeast-1:111111111111:components:component-name
output "greengrass_component" {
  value = provider::arn::greengrass_component("component-name")
}
