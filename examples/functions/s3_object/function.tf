# arn:aws:s3:::bucket-name/object-name
output "s3_object" {
  value = provider::arn::s3_object("bucket-name", "object-name")
}
