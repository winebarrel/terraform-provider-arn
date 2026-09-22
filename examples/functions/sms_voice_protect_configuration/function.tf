# arn:aws:sms-voice:ap-northeast-1:111111111111:protect-configuration/protect-configuration-id
output "sms_voice_protect_configuration" {
  value = provider::arn::sms_voice_protect_configuration("protect-configuration-id")
}
