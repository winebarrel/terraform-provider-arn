# arn:aws:s3-outposts:ap-northeast-1:111111111111:outpost/outpost-id/endpoint/endpoint-id
output "s3_outposts_endpoint" {
  value = provider::arn::s3_outposts_endpoint("outpost-id", "endpoint-id")
}
