# arn:aws:events:ap-northeast-1:111111111111:api-destination/api-destination-name
output "events_api_destination" {
  value = provider::arn::events_api_destination("api-destination-name")
}
