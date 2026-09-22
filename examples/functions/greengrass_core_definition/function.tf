# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/cores/core-definition-id
output "greengrass_core_definition" {
  value = provider::arn::greengrass_core_definition("core-definition-id")
}
