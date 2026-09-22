# arn:aws:rds:ap-northeast-1:111111111111:auto-backup:db-instance-automated-backup-id
output "rds_auto_backup" {
  value = provider::arn::rds_auto_backup("db-instance-automated-backup-id")
}
