# arn:aws:ivs:ap-northeast-1:111111111111:stream-key/resource-id
output "ivs_stream_key" {
  value = provider::arn::ivs_stream_key("resource-id")
}
