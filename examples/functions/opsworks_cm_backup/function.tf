# arn:aws:opsworks-cm::111111111111:backup/server-name-{Date-and-Time-Stamp-of-Backup}
output "opsworks_cm_backup" {
  value = provider::arn::opsworks_cm_backup("server-name")
}
