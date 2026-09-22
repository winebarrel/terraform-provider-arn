# arn:aws:kinesis:ap-northeast-1:111111111111:stream-type/stream-name/consumer/consumer-name:consumer-creation-timpstamp
output "kinesis_consumer" {
  value = provider::arn::kinesis_consumer("stream-type", "stream-name", "consumer-name", "consumer-creation-timpstamp")
}
