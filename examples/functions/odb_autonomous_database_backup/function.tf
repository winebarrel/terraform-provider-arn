# arn:aws:odb:ap-northeast-1:111111111111:autonomous-database-backup/autonomous-database-backup-id
output "odb_autonomous_database_backup" {
  value = provider::arn::odb_autonomous_database_backup("autonomous-database-backup-id")
}
