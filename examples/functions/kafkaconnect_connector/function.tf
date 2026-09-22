# arn:aws:kafkaconnect:ap-northeast-1:111111111111:connector/connector-name/uuid
output "kafkaconnect_connector" {
  value = provider::arn::kafkaconnect_connector("connector-name", "uuid")
}
