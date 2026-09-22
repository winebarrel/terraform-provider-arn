# arn:aws:s3-outposts:ap-northeast-1:111111111111:outpost/outpost-id/accesspoint/access-point-name
output "s3_outposts_accesspoint" {
  value = provider::arn::s3_outposts_accesspoint("outpost-id", "access-point-name")
}
