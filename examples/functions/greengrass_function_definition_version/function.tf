# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/functions/function-definition-id/versions/version-id
output "greengrass_function_definition_version" {
  value = provider::arn::greengrass_function_definition_version("function-definition-id", "version-id")
}
