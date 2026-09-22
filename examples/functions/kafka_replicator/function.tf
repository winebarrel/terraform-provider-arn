# arn:aws:kafka:ap-northeast-1:111111111111:replicator/replicator-name/uuid
output "kafka_replicator" {
  value = provider::arn::kafka_replicator("replicator-name", "uuid")
}
