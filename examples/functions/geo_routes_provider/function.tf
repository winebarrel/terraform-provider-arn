# arn:aws:geo-routes:ap-northeast-1::provider/default
output "geo_routes_provider" {
  value = provider::arn::geo_routes_provider()
}
