# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/loggers/logger-definition-id/versions/version-id
output "greengrass_logger_definition_version" {
  value = provider::arn::greengrass_logger_definition_version("logger-definition-id", "version-id")
}
