# arn:aws:s3:::my-bucket
#
# An S3 bucket ARN carries neither an account nor a region, so this one works
# with no configuration file at all.
output "bucket" {
  value = provider::arn::s3_bucket("my-bucket")
}
