# arn:aws:s3-outposts:ap-northeast-1:111111111111:outpost/outpost-id/bucket/bucket-name/object/object-name
output "s3_outposts_object" {
  value = provider::arn::s3_outposts_object("outpost-id", "bucket-name", "object-name")
}
