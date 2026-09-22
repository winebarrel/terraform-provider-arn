# arn:aws:mediapackagev2:ap-northeast-1:111111111111:channelGroup/channel-group-name
output "mediapackagev2_channel_group" {
  value = provider::arn::mediapackagev2_channel_group("channel-group-name")
}
