# arn:aws:medialive:ap-northeast-1:111111111111:channel:channel-id
output "medialive_channel" {
  value = provider::arn::medialive_channel("channel-id")
}
