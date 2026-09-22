# arn:aws:sms-voice:ap-northeast-1:111111111111:registration/registration-id
output "sms_voice_registration" {
  value = provider::arn::sms_voice_registration("registration-id")
}
