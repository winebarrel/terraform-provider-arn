# arn:aws:s3-outposts:ap-northeast-1:111111111111:outpost/outpost-id/bucket/bucket-name
output "s3_outposts_bucket" {
  value = provider::arn::s3_outposts_bucket("outpost-id", "bucket-name")
}
