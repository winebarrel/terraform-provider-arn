# arn:aws:kinesisvideo:ap-northeast-1:111111111111:channel/channel-name/creation-time
output "kinesisvideo_channel" {
  value = provider::arn::kinesisvideo_channel("channel-name", "creation-time")
}
