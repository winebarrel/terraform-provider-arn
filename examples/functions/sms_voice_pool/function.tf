# arn:aws:sms-voice:ap-northeast-1:111111111111:pool/pool-id
output "sms_voice_pool" {
  value = provider::arn::sms_voice_pool("pool-id")
}
