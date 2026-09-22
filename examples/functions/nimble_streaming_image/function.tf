# arn:aws:nimble:ap-northeast-1:111111111111:streaming-image/streaming-image-id
output "nimble_streaming_image" {
  value = provider::arn::nimble_streaming_image("streaming-image-id")
}
