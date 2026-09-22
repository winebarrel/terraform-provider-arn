# arn:aws:s3:::bucket-name
output "s3_bucket" {
  value = provider::arn::s3_bucket("bucket-name")
}
