# arn:aws:chime:ap-northeast-1:111111111111:voice-profile/voice-profile-id
output "chime_voice_profile" {
  value = provider::arn::chime_voice_profile("voice-profile-id")
}
