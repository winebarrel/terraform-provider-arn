# arn:aws:datasync:ap-northeast-1:111111111111:location/location-id
output "datasync_location" {
  value = provider::arn::datasync_location("location-id")
}
