# arn:aws:chime:ap-northeast-1:111111111111:media-pipeline-kinesis-video-stream-pool/pool-name
output "chime_media_pipeline_kinesis_video_stream_pool" {
  value = provider::arn::chime_media_pipeline_kinesis_video_stream_pool("pool-name")
}
