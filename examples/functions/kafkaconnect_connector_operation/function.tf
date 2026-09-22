# arn:aws:kafkaconnect:ap-northeast-1:111111111111:connector-operation/connector-name/connector-uuid/uuid
output "kafkaconnect_connector_operation" {
  value = provider::arn::kafkaconnect_connector_operation("connector-name", "connector-uuid", "uuid")
}
