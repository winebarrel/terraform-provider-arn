# arn:aws:backup:ap-northeast-1:111111111111:backup-vault:backup-vault-name
output "backup_backup_vault" {
  value = provider::arn::backup_backup_vault("backup-vault-name")
}
