# arn:aws:config:ap-northeast-1:111111111111:connector/provider/provider-id/connector-id
output "config_connector" {
  value = provider::arn::config_connector("provider", "provider-id", "connector-id")
}
