# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/cores/core-definition-id/versions/version-id
output "greengrass_core_definition_version" {
  value = provider::arn::greengrass_core_definition_version("core-definition-id", "version-id")
}
