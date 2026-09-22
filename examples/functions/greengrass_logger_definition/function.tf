# arn:aws:greengrass:ap-northeast-1:111111111111:/greengrass/definition/loggers/logger-definition-id
output "greengrass_logger_definition" {
  value = provider::arn::greengrass_logger_definition("logger-definition-id")
}
