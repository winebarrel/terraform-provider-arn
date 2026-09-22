# arn:aws:iot:ap-northeast-1:111111111111:stream/stream-id
output "iot_stream" {
  value = provider::arn::iot_stream("stream-id")
}
