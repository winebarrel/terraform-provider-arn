# arn:aws:kafka:ap-northeast-1:111111111111:transactional-id/cluster-name/cluster-uuid/transactional-id
output "kafka_transactional_id" {
  value = provider::arn::kafka_transactional_id("cluster-name", "cluster-uuid", "transactional-id")
}
