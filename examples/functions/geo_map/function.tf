# arn:aws:geo:ap-northeast-1:111111111111:map/map-name
output "geo_map" {
  value = provider::arn::geo_map("map-name")
}
