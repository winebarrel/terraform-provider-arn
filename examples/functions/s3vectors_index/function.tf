# arn:aws:s3vectors:ap-northeast-1:111111111111:bucket/bucket-name/index/index-name
output "s3vectors_index" {
  value = provider::arn::s3vectors_index("bucket-name", "index-name")
}
