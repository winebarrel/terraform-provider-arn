# arn:aws:chime:ap-northeast-1:111111111111:media-pipeline/media-pipeline-id
output "chime_media_pipeline" {
  value = provider::arn::chime_media_pipeline("media-pipeline-id")
}
