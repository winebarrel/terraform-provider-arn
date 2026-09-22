# arn:aws:s3files:ap-northeast-1:111111111111:file-system/file-system-id/access-point/access-point-id
output "s3files_access_point" {
  value = provider::arn::s3files_access_point("file-system-id", "access-point-id")
}
