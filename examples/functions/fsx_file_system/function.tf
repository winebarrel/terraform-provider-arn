# arn:aws:fsx:ap-northeast-1:111111111111:file-system/file-system-id
output "fsx_file_system" {
  value = provider::arn::fsx_file_system("file-system-id")
}
