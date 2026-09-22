# arn:aws:sms-voice:ap-northeast-1:111111111111:configuration-set/configuration-set-name
output "sms_voice_configuration_set" {
  value = provider::arn::sms_voice_configuration_set("configuration-set-name")
}
