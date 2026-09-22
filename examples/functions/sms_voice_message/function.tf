# arn:aws:sms-voice:ap-northeast-1:111111111111:message/message-id
output "sms_voice_message" {
  value = provider::arn::sms_voice_message("message-id")
}
