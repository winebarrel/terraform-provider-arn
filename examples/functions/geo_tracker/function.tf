# arn:aws:geo:ap-northeast-1:111111111111:tracker/tracker-name
output "geo_tracker" {
  value = provider::arn::geo_tracker("tracker-name")
}
