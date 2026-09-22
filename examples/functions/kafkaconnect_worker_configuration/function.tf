# arn:aws:kafkaconnect:ap-northeast-1:111111111111:worker-configuration/worker-configuration-name/uuid
output "kafkaconnect_worker_configuration" {
  value = provider::arn::kafkaconnect_worker_configuration("worker-configuration-name", "uuid")
}
