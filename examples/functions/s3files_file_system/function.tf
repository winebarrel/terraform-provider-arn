# arn:aws:s3files:ap-northeast-1:111111111111:file-system/file-system-id
output "s3files_file_system" {
  value = provider::arn::s3files_file_system("file-system-id")
}
