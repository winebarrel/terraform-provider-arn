# arn:aws:geo-places:ap-northeast-1::provider/default
output "geo_places_provider" {
  value = provider::arn::geo_places_provider()
}
