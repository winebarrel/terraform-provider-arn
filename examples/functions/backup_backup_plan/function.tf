# arn:aws:backup:ap-northeast-1:111111111111:backup-plan:backup-plan-id
output "backup_backup_plan" {
  value = provider::arn::backup_backup_plan("backup-plan-id")
}
