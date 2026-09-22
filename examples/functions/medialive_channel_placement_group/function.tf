# arn:aws:medialive:ap-northeast-1:111111111111:channelPlacementGroup:cluster-id/channel-placement-group-id
output "medialive_channel_placement_group" {
  value = provider::arn::medialive_channel_placement_group("cluster-id", "channel-placement-group-id")
}
