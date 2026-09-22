# arn:aws:mediatailor:ap-northeast-1:111111111111:sourceLocation/source-location-name
output "mediatailor_source_location" {
  value = provider::arn::mediatailor_source_location("source-location-name")
}
