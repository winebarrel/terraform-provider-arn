# arn:aws:mediapackage:ap-northeast-1:111111111111:origin_endpoints/origin-endpoint-identifier
output "mediapackage_origin_endpoints" {
  value = provider::arn::mediapackage_origin_endpoints("origin-endpoint-identifier")
}
