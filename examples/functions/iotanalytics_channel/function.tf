# arn:aws:iotanalytics:ap-northeast-1:111111111111:channel/channel-name
output "iotanalytics_channel" {
  value = provider::arn::iotanalytics_channel("channel-name")
}
