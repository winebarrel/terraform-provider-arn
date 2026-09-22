# arn:aws:kinesisvideo:ap-northeast-1:111111111111:stream/stream-name/creation-time
output "kinesisvideo_stream" {
  value = provider::arn::kinesisvideo_stream("stream-name", "creation-time")
}
