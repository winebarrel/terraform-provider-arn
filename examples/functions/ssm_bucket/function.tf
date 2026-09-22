# arn:aws:s3:::bucket-name
output "ssm_bucket" {
  value = provider::arn::ssm_bucket("bucket-name")
}
