# arn:aws:elasticfilesystem:ap-northeast-1:111111111111:file-system/file-system-id
output "elasticfilesystem_file_system" {
  value = provider::arn::elasticfilesystem_file_system("file-system-id")
}
