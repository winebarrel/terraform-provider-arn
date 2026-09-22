# arn:aws:sms-voice:ap-northeast-1:111111111111:phone-number/phone-number-id
output "sms_voice_phone_number" {
  value = provider::arn::sms_voice_phone_number("phone-number-id")
}
