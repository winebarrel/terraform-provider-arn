# arn:aws:sms-voice:ap-northeast-1:111111111111:sender-id/sender-id/iso-country-code
output "sms_voice_sender_id" {
  value = provider::arn::sms_voice_sender_id("sender-id", "iso-country-code")
}
