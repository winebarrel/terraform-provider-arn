# arn:aws:mediapackagev2:ap-northeast-1:111111111111:channelGroup/channel-group-name/channel/channel-name/originEndpoint/origin-endpoint-name
output "mediapackagev2_origin_endpoint" {
  value = provider::arn::mediapackagev2_origin_endpoint("channel-group-name", "channel-name", "origin-endpoint-name")
}
