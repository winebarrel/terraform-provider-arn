# arn:aws:sms-voice:ap-northeast-1:111111111111:notify-configuration/notify-configuration-id
output "sms_voice_notify_configuration" {
  value = provider::arn::sms_voice_notify_configuration("notify-configuration-id")
}
