# arn:aws:mediatailor:ap-northeast-1:111111111111:liveSource/source-location-name/live-source-name
output "mediatailor_live_source" {
  value = provider::arn::mediatailor_live_source("source-location-name", "live-source-name")
}
