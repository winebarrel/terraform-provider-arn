# arn:aws:dms:ap-northeast-1:111111111111:task:*
output "dms_replication_task" {
  value = provider::arn::dms_replication_task()
}
