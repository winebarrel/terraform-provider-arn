# arn:aws:chime:ap-northeast-1:111111111111:app-instance/app-instance-id/channel-flow/channel-flow-id
output "chime_channel_flow" {
  value = provider::arn::chime_channel_flow("app-instance-id", "channel-flow-id")
}
