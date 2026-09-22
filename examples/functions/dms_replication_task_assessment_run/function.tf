# arn:aws:dms:ap-northeast-1:111111111111:assessment-run:*
output "dms_replication_task_assessment_run" {
  value = provider::arn::dms_replication_task_assessment_run()
}
