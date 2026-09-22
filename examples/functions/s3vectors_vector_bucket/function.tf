# arn:aws:s3vectors:ap-northeast-1:111111111111:bucket/bucket-name
output "s3vectors_vector_bucket" {
  value = provider::arn::s3vectors_vector_bucket("bucket-name")
}
