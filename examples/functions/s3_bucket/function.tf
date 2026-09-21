# arn:aws:s3:::my-bucket
output "bucket" {
  value = provider::arn::s3_bucket("my-bucket")
}
