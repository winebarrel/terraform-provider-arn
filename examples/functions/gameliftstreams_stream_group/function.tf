# arn:aws:gameliftstreams:ap-northeast-1:111111111111:streamgroup/stream-group-id
output "gameliftstreams_stream_group" {
  value = provider::arn::gameliftstreams_stream_group("stream-group-id")
}
