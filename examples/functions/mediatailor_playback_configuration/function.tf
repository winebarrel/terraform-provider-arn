# arn:aws:mediatailor:ap-northeast-1:111111111111:playbackConfiguration/resource-id
output "mediatailor_playback_configuration" {
  value = provider::arn::mediatailor_playback_configuration("resource-id")
}
