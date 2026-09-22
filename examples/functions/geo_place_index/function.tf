# arn:aws:geo:ap-northeast-1:111111111111:place-index/index-name
output "geo_place_index" {
  value = provider::arn::geo_place_index("index-name")
}
