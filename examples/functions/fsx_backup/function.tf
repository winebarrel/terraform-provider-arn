# arn:aws:fsx:ap-northeast-1:111111111111:backup/backup-id
output "fsx_backup" {
  value = provider::arn::fsx_backup("backup-id")
}
