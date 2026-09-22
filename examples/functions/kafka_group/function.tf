# arn:aws:kafka:ap-northeast-1:111111111111:group/cluster-name/cluster-uuid/group-name
output "kafka_group" {
  value = provider::arn::kafka_group("cluster-name", "cluster-uuid", "group-name")
}
