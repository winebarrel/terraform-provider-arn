# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/functions/function-definition-id
output "greengrass_function_definition" {
  value = provider::arn::greengrass_function_definition("function-definition-id")
}
