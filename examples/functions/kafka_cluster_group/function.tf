# arn:aws:kafka:ap-northeast-1:111111111111:group/cluster-name/cluster-uuid/group-name
output "kafka_cluster_group" {
  value = provider::arn::kafka_cluster_group("cluster-name", "cluster-uuid", "group-name")
}
