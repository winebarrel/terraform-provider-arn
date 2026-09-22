# arn:aws:chime:ap-northeast-1:111111111111:vc/voice-connector-id
output "chime_voice_connector" {
  value = provider::arn::chime_voice_connector("voice-connector-id")
}
