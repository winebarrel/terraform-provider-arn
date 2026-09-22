# arn:aws:kafka:ap-northeast-1:111111111111:vpc-connection/cluster-owner-account/cluster-name/uuid
output "kafka_vpc_connection" {
  value = provider::arn::kafka_vpc_connection("cluster-owner-account", "cluster-name", "uuid")
}
