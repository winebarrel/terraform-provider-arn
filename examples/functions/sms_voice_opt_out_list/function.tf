# arn:aws:sms-voice:ap-northeast-1:111111111111:opt-out-list/opt-out-list-name
output "sms_voice_opt_out_list" {
  value = provider::arn::sms_voice_opt_out_list("opt-out-list-name")
}
