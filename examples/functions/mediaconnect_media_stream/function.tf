# arn:aws:mediaconnect:ap-northeast-1:111111111111:flow:flow-id:flow-name/mediaStream/media-stream-name
output "mediaconnect_media_stream" {
  value = provider::arn::mediaconnect_media_stream("flow-id", "flow-name", "media-stream-name")
}
