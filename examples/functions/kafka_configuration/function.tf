# arn:aws:kafka:ap-northeast-1:111111111111:configuration/configuration-name/uuid
output "kafka_configuration" {
  value = provider::arn::kafka_configuration("configuration-name", "uuid")
}
