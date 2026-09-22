# arn:aws:mgh:ap-northeast-1:111111111111:progressUpdateStream/stream
output "mgh_progress_update_stream" {
  value = provider::arn::mgh_progress_update_stream("stream")
}
