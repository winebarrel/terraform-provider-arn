# arn:aws:kinesis:ap-northeast-1:111111111111:stream/stream-name
output "kinesis_stream" {
  value = provider::arn::kinesis_stream("stream-name")
}
