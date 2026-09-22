# arn:aws:chime:ap-northeast-1:111111111111:voice-profile-domain/voice-profile-domain-id
output "chime_voice_profile_domain" {
  value = provider::arn::chime_voice_profile_domain("voice-profile-domain-id")
}
