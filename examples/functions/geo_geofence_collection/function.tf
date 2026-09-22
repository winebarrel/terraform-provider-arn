# arn:aws:geo:ap-northeast-1:111111111111:geofence-collection/geofence-collection-name
output "geo_geofence_collection" {
  value = provider::arn::geo_geofence_collection("geofence-collection-name")
}
