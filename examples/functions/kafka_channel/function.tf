# arn:aws:kafka:ap-northeast-1:111111111111:channel/cluster-name/cluster-uuid/channel-name/uuid
output "kafka_channel" {
  value = provider::arn::kafka_channel("cluster-name", "cluster-uuid", "channel-name", "uuid")
}
