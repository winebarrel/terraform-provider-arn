# arn:aws:nimble:ap-northeast-1:111111111111:streaming-session/streaming-session-id
output "nimble_streaming_session" {
  value = provider::arn::nimble_streaming_session("streaming-session-id")
}
