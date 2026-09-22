# arn:aws:kafka:ap-northeast-1:111111111111:topic/cluster-name/cluster-uuid/topic-name
output "kafka_cluster_topic" {
  value = provider::arn::kafka_cluster_topic("cluster-name", "cluster-uuid", "topic-name")
}
