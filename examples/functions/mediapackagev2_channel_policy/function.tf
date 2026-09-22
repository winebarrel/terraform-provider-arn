# arn:aws:mediapackagev2:ap-northeast-1:111111111111:channelGroup/channel-group-name/channel/channel-name
output "mediapackagev2_channel_policy" {
  value = provider::arn::mediapackagev2_channel_policy("channel-group-name", "channel-name")
}
