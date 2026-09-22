# arn:aws:kafka:ap-northeast-1:111111111111:cluster/cluster-name/uuid
output "kafka_cluster" {
  value = provider::arn::kafka_cluster("cluster-name", "uuid")
}
