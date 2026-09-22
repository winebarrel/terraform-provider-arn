# arn:aws:fsx:ap-northeast-1:111111111111:file-cache/file-cache-id
output "fsx_file_cache" {
  value = provider::arn::fsx_file_cache("file-cache-id")
}
