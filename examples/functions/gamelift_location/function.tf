# arn:aws:gamelift:ap-northeast-1:111111111111:location/location-id
output "gamelift_location" {
  value = provider::arn::gamelift_location("location-id")
}
