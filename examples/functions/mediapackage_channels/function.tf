# arn:aws:mediapackage:ap-northeast-1:111111111111:channels/channel-identifier
output "mediapackage_channels" {
  value = provider::arn::mediapackage_channels("channel-identifier")
}
