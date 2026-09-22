# arn:aws:chime:ap-northeast-1:111111111111:app-instance/app-instance-id/channel/channel-id
output "chime_channel" {
  value = provider::arn::chime_channel("app-instance-id", "channel-id")
}
