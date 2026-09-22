# arn:aws:sms-voice:ap-northeast-1:111111111111:verified-destination-number/verified-destination-number-id
output "sms_voice_verified_destination_number" {
  value = provider::arn::sms_voice_verified_destination_number("verified-destination-number-id")
}
