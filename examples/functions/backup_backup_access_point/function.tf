# arn:aws:backup:ap-northeast-1:111111111111:accesspoint/access-point-name
output "backup_backup_access_point" {
  value = provider::arn::backup_backup_access_point("access-point-name")
}
