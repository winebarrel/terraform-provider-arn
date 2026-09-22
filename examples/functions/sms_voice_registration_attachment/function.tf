# arn:aws:sms-voice:ap-northeast-1:111111111111:registration-attachment/registration-attachment-id
output "sms_voice_registration_attachment" {
  value = provider::arn::sms_voice_registration_attachment("registration-attachment-id")
}
