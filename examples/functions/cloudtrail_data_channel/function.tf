# arn:aws:cloudtrail:ap-northeast-1:111111111111:channel/channel-id
output "cloudtrail_data_channel" {
  value = provider::arn::cloudtrail_data_channel("channel-id")
}
