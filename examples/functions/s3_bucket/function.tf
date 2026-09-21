# arn:aws:s3:::my-bucket
#
# An S3 bucket ARN carries neither an account nor a region, so an empty
# .arn.hcl is enough for this one. The file still has to exist.
output "bucket" {
  value = provider::arn::s3_bucket("my-bucket")
}
