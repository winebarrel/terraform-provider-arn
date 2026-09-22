# arn:aws:dms:ap-northeast-1:111111111111:replication-config:*
output "dms_replication_config" {
  value = provider::arn::dms_replication_config()
}
