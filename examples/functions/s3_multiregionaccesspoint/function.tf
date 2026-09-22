# arn:aws:s3::111111111111:accesspoint/access-point-alias
output "s3_multiregionaccesspoint" {
  value = provider::arn::s3_multiregionaccesspoint("access-point-alias")
}
