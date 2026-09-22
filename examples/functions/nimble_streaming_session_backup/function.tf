# arn:aws:nimble:ap-northeast-1:111111111111:streaming-session-backup/streaming-session-backup-id
output "nimble_streaming_session_backup" {
  value = provider::arn::nimble_streaming_session_backup("streaming-session-backup-id")
}
