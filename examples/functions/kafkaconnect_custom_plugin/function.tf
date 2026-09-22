# arn:aws:kafkaconnect:ap-northeast-1:111111111111:custom-plugin/custom-plugin-name/uuid
output "kafkaconnect_custom_plugin" {
  value = provider::arn::kafkaconnect_custom_plugin("custom-plugin-name", "uuid")
}
