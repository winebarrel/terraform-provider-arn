# arn:aws:geo:ap-northeast-1:111111111111:api-key/key-name
output "geo_api_key" {
  value = provider::arn::geo_api_key("key-name")
}
