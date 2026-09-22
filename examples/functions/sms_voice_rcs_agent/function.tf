# arn:aws:sms-voice:ap-northeast-1:111111111111:rcs-agent/rcs-agent-id
output "sms_voice_rcs_agent" {
  value = provider::arn::sms_voice_rcs_agent("rcs-agent-id")
}
