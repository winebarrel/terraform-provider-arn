# arn:aws:geo-maps:ap-northeast-1::provider/default
output "geo_maps_provider" {
  value = provider::arn::geo_maps_provider()
}
