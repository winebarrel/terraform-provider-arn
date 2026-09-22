# arn:aws:dynamodb:ap-northeast-1:111111111111:table/table-name/backup/backup-name
output "dynamodb_backup" {
  value = provider::arn::dynamodb_backup("table-name", "backup-name")
}
