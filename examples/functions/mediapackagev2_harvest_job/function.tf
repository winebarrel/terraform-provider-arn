# arn:aws:mediapackagev2:ap-northeast-1:111111111111:channelGroup/channel-group-name/channel/channel-name/originEndpoint/origin-endpoint-name/harvestJob/harvest-job-name
output "mediapackagev2_harvest_job" {
  value = provider::arn::mediapackagev2_harvest_job("channel-group-name", "channel-name", "origin-endpoint-name", "harvest-job-name")
}
