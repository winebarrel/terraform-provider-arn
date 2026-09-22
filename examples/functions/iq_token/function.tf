# arn:aws:iq:ap-northeast-1::token/token-id
output "iq_token" {
  value = provider::arn::iq_token("token-id")
}
