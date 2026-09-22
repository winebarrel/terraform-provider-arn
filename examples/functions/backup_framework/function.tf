# arn:aws:backup:ap-northeast-1:111111111111:framework:framework-name-framework-id
output "backup_framework" {
  value = provider::arn::backup_framework("framework-name", "framework-id")
}
