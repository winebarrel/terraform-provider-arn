# arn:aws:mobiletargeting:ap-northeast-1:111111111111:apps/app-id/channels/channel-type
output "mobiletargeting_channel" {
  value = provider::arn::mobiletargeting_channel("app-id", "channel-type")
}
