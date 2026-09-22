# arn:aws:mediatailor:ap-northeast-1:111111111111:channel/channel-name
output "mediatailor_channel" {
  value = provider::arn::mediatailor_channel("channel-name")
}
