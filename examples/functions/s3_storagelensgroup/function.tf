# arn:aws:s3:ap-northeast-1:111111111111:storage-lens-group/name
output "s3_storagelensgroup" {
  value = provider::arn::s3_storagelensgroup("name")
}
