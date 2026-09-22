# arn:aws:dms:ap-northeast-1:111111111111:rep:*
output "dms_replication_instance" {
  value = provider::arn::dms_replication_instance()
}
