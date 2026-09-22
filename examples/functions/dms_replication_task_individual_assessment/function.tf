# arn:aws:dms:ap-northeast-1:111111111111:individual-assessment:*
output "dms_replication_task_individual_assessment" {
  value = provider::arn::dms_replication_task_individual_assessment()
}
