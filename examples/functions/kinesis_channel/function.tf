# arn:aws:kinesis:ap-northeast-1:111111111111:channel/channel-id
output "kinesis_channel" {
  value = provider::arn::kinesis_channel("channel-id")
}
