# arn:aws:ivschat:ap-northeast-1:111111111111:logging-configuration/resource-id
output "ivschat_logging_configuration" {
  value = provider::arn::ivschat_logging_configuration("resource-id")
}
