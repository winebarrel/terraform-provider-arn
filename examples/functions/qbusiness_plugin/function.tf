# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/plugin/plugin-id
output "qbusiness_plugin" {
  value = provider::arn::qbusiness_plugin("application-id", "plugin-id")
}
