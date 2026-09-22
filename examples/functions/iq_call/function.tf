# arn:aws:iq:ap-northeast-1::call/call-id
output "iq_call" {
  value = provider::arn::iq_call("call-id")
}
